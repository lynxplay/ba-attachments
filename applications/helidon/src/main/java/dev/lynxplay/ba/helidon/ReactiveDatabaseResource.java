package dev.lynxplay.ba.helidon;

import io.helidon.common.reactive.Single;
import io.helidon.dbclient.DbClient;
import io.helidon.dbclient.DbRow;

import java.math.BigDecimal;
import java.util.UUID;

public class ReactiveDatabaseResource {

    public Single<ProductRecord> findRecord(DbClient databaseClient, UUID recordIdentifier) {
        return databaseClient.execute(e -> e.createQuery("SELECT identifier, price, title, description FROM products WHERE identifier = ?")
            .params(recordIdentifier)
            .execute()
            .first()
            .map(this::parseRecord));
    }

    private ProductRecord parseRecord(DbRow row) {
        return new ProductRecord(
            row.column("identifier").as(UUID.class),
            row.column("price").as(BigDecimal.class),
            row.column("title").as(String.class),
            row.column("description").as(String.class)
        );
    }
}
