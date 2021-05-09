package dev.lynxplay.ba.quarkus;

import io.smallrye.mutiny.Uni;
import io.vertx.mutiny.pgclient.PgPool;
import io.vertx.mutiny.sqlclient.Row;
import io.vertx.mutiny.sqlclient.RowSet;
import io.vertx.mutiny.sqlclient.Tuple;

import javax.enterprise.context.ApplicationScoped;
import java.util.UUID;

@ApplicationScoped
public class ReactiveDatabaseResource {

    public Uni<ProductRecord> findRecord(PgPool databaseClient, UUID recordIdentifier) {
        return databaseClient.preparedQuery("SELECT identifier, price, title, description FROM products WHERE identifier = $1")
            .execute(Tuple.of(recordIdentifier))
            .onItem().transform(RowSet::iterator)
            .onItem().transform(i -> i.hasNext() ? parseRecord(i.next()) : null)
            .onItem().ifNull().failWith(new RuntimeException("Could not find product record " + recordIdentifier + "!"));
    }

    private ProductRecord parseRecord(Row row) {
        return new ProductRecord(
            row.getUUID("identifier"),
            row.getBigDecimal("price"),
            row.getString("title"),
            row.getString("description")
        );
    }
}
