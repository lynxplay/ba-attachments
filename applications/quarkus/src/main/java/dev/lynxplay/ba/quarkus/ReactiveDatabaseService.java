package dev.lynxplay.ba.quarkus;

import io.quarkus.vertx.web.Route;
import io.smallrye.mutiny.Uni;
import io.vertx.core.http.HttpMethod;
import io.vertx.ext.web.RoutingContext;
import io.vertx.mutiny.pgclient.PgPool;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import java.util.UUID;

@ApplicationScoped
public class ReactiveDatabaseService {

    private final ReactiveDatabaseResource databaseResource;
    private final PgPool databaseClient;

    @Inject
    public ReactiveDatabaseService(ReactiveDatabaseResource databaseResource, PgPool databaseClient) {
        this.databaseResource = databaseResource;
        this.databaseClient = databaseClient;
    }

    @Route(path = "/product/:identifier", methods = HttpMethod.GET, produces = "application/json")
    public Uni<ProductRecord> product(RoutingContext context) {
        return databaseResource.findRecord(this.databaseClient, UUID.fromString(context.request().getParam("identifier")));
    }

}
