package dev.lynxplay.ba.micronaut;

import io.reactivex.Maybe;
import io.vertx.reactivex.pgclient.PgPool;
import io.vertx.reactivex.sqlclient.Row;
import io.vertx.reactivex.sqlclient.RowSet;
import io.vertx.reactivex.sqlclient.Tuple;

import javax.inject.Singleton;
import java.util.UUID;

@Singleton
public class ReactiveDatabaseResource {

    public Maybe<ProductRecord> findRecord(PgPool databaseClient, UUID recordIdentifier) {
        return databaseClient.preparedQuery("SELECT identifier, price, title, description FROM products WHERE identifier = $1")
            .rxExecute(Tuple.of(recordIdentifier))
            .map(RowSet::iterator)
            .flatMapMaybe(i -> i.hasNext() ? Maybe.just(parseRecord(i.next())) : Maybe.empty());
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
