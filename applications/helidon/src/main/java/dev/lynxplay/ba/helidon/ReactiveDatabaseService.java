package dev.lynxplay.ba.helidon;

import io.helidon.dbclient.DbClient;
import io.helidon.webserver.Routing;
import io.helidon.webserver.ServerRequest;
import io.helidon.webserver.ServerResponse;
import io.helidon.webserver.Service;

import java.util.UUID;

public class ReactiveDatabaseService implements Service {

    private final ReactiveDatabaseResource databaseResource;
    private final DbClient databaseClient;

    public ReactiveDatabaseService(ReactiveDatabaseResource databaseResource, DbClient databaseClient) {
        this.databaseResource = databaseResource;
        this.databaseClient = databaseClient;
    }

    @Override
    public void update(Routing.Rules rules) {
        rules.get("/{identifier}", this::handleProductGet);
    }

    public void handleProductGet(ServerRequest request, ServerResponse response) {
        final UUID identifier = UUID.fromString(request.path().param("identifier"));
        this.databaseResource.findRecord(this.databaseClient, identifier).subscribe(response::send, throwable -> {
            throwable.printStackTrace();
            response.status(500).send();
        });
    }
}
