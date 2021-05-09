package dev.lynxplay.ba.springboot;

import io.r2dbc.spi.Row;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.r2dbc.core.DatabaseClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;
import reactor.util.annotation.NonNull;

import java.math.BigDecimal;
import java.util.UUID;

@RestController
@RequestMapping("/product")
public class ReactiveProductHandler {

    private final DatabaseClient databaseClient;

    @Autowired
    public ReactiveProductHandler(DatabaseClient databaseClient) {
        this.databaseClient = databaseClient;
    }

    @NonNull
    @GetMapping("/{identifier}")
    public Mono<ProductRecord> getProduct(@PathVariable UUID identifier) {
        return this.databaseClient.sql("SELECT identifier, price, title, description FROM products WHERE identifier = $1")
            .bind(0, identifier)
            .map(this::parseRecord)
            .first();
    }

    private ProductRecord parseRecord(Row row) {
        return new ProductRecord(
            row.get("identifier", UUID.class),
            row.get("price", BigDecimal.class),
            row.get("title", String.class),
            row.get("description", String.class)
        );
    }
}
