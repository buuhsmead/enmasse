/**
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 * <p>
 * http://www.apache.org/licenses/LICENSE-2.0
 * <p>
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.enmasse.keycloak.spi;

import io.vertx.core.AbstractVerticle;
import io.vertx.core.net.JksOptions;
import io.vertx.core.net.PemKeyCertOptions;
import io.vertx.core.net.PfxOptions;
import io.vertx.proton.ProtonConnection;
import io.vertx.proton.ProtonServer;
import io.vertx.proton.ProtonServerOptions;
import org.apache.qpid.proton.amqp.Symbol;
import org.jboss.logging.Logger;
import org.keycloak.Config;
import org.keycloak.models.GroupModel;
import org.keycloak.models.KeycloakSessionFactory;
import org.keycloak.models.UserModel;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

public class AmqpServer extends AbstractVerticle {

    private static final Logger LOG = Logger.getLogger(AmqpServer.class);

    private final String hostname;
    private final int port;
    private final Config.Scope config;
    private final boolean useTls;
    private volatile ProtonServer server;
    private KeycloakSessionFactory keycloakSessionFactory;

    public AmqpServer(String hostname, int port, final Config.Scope config, final boolean useTls) {
        this.hostname = hostname;
        this.port = port;
        this.config = config;
        this.useTls = useTls;
    }

    private void connectHandler(ProtonConnection connection) {
        String containerId = config.get("containerId", "keycloak-amqp");
        connection.setContainer(containerId);
        connection.openHandler(conn -> {
            UserModel userModel = connection.attachments().get(SaslAuthenticator.USER_ATTACHMENT, UserModel.class);
            if(userModel != null) {
                Map<Symbol, Object> props = new HashMap<>();
                Map<String, Object> authUserMap = new HashMap<>();
                authUserMap.put("sub", userModel.getId());
                authUserMap.put("preferred_username", userModel.getUsername());
                props.put(Symbol.valueOf("authenticated-identity"), authUserMap);
                Set<String> groups = userModel.getGroups().stream().map(GroupModel::getName).collect(Collectors.toSet());
                props.put(Symbol.valueOf("groups"), new ArrayList<>(groups));
                connection.setProperties(props);
            }
            connection.open();
            connection.close();
        }).closeHandler(conn -> {
            connection.close();
            connection.disconnect();
        }).disconnectHandler(protonConnection -> {
            connection.disconnect();
        });

    }

    @Override
    public void start() {
        ProtonServerOptions options = new ProtonServerOptions();
        if(useTls) {
            options.setSsl(true);
            String path;
            if((path = config.get("jksKeyStorePath")) != null) {
                final JksOptions jksOptions = new JksOptions();
                jksOptions.setPath(path);
                jksOptions.setPassword(config.get("keyStorePassword"));
                options.setKeyStoreOptions(jksOptions);
            } else if((path = config.get("pfxKeyStorePath")) != null) {
                final PfxOptions pfxOptions = new PfxOptions();
                pfxOptions.setPath(path);
                pfxOptions.setPassword(config.get("keyStorePassword"));
                options.setPfxKeyCertOptions(pfxOptions);
            } else if((path = config.get("pemCertificatePath")) != null) {
                final PemKeyCertOptions pemKeyCertOptions = new PemKeyCertOptions();
                pemKeyCertOptions.setCertPath(path);
                pemKeyCertOptions.setKeyPath(config.get("pemKeyPath"));
                options.setPemKeyCertOptions(pemKeyCertOptions);
            } else {
                // use JDK settings?
            }

        }
        server = ProtonServer.create(vertx, options);

        server.saslAuthenticatorFactory(() -> new SaslAuthenticator(keycloakSessionFactory, config, useTls));
        server.connectHandler(this::connectHandler);
        LOG.info("Starting server on "+hostname+":"+ port);
        server.listen(port, hostname, event -> {
            if(event.failed())
            {
                LOG.error("Unable to listen for AMQP on "+hostname+":" + port, event.cause());
            }

        });
    }

    @Override
    public void stop() {
        if (server != null) {
            server.close();
        }
    }

    void setKeycloakSessionFactory(final KeycloakSessionFactory keycloakSessionFactory)
    {
        this.keycloakSessionFactory = keycloakSessionFactory;
    }
}
