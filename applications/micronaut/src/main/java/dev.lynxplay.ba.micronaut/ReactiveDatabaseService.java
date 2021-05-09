package dev.lynxplay.ba.micronaut;

import io.micronaut.http.MediaType;
import io.micronaut.http.annotation.Controller;
import io.micronaut.http.annotation.Get;
import io.micronaut.http.annotation.PathVariable;
import io.reactivex.Maybe;
import io.vertx.reactivex.pgclient.PgPool;

import javax.inject.Inject;
import java.util.UUID;

@Controller
public class ReactiveDatabaseService {

    private final ReactiveDatabaseResource databaseResource;
    private final PgPool databaseClient;

    @Inject
    public ReactiveDatabaseService(ReactiveDatabaseResource databaseResource, PgPool databaseClient) {
        this.databaseResource = databaseResource;
        this.databaseClient = databaseClient;
    }

    @Get(value = "/product/{identifier}", produces = MediaType.APPLICATION_JSON)
    public Maybe<ProductRecord> product(@PathVariable UUID identifier) {
        return databaseResource.findRecord(this.databaseClient, identifier);
    }
}
