
package dev.lynxplay.ba.helidon;

import io.helidon.common.LogConfig;
import io.helidon.config.Config;
import io.helidon.dbclient.DbClient;
import io.helidon.media.jackson.JacksonSupport;
import io.helidon.webserver.Routing;
import io.helidon.webserver.WebServer;

/**
 * The application main class.
 */
public final class Application {

    public static void main(final String[] args) {
        LogConfig.configureRuntime();
        final Config config = Config.create();

        final DbClient databaseClient = DbClient.create(config.get("db"));

        final Routing router = Routing.builder()
            .register("/product", new ReactiveDatabaseService(new ReactiveDatabaseResource(), databaseClient))
            .build();

        final WebServer server = WebServer.builder()
            .config(config.get("server"))
            .workersCount(25)
            .addMediaSupport(JacksonSupport.create())
            .routing(router)
            .build();

        server.start().thenAccept(ws -> {
            System.out.println("Server is running on port " + ws.port());
            ws.whenShutdown().thenRun(() -> System.out.println("Shutting down!"));
        });
    }
}
