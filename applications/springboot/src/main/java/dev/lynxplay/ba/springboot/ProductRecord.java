package dev.lynxplay.ba.springboot;

import org.springframework.data.annotation.Id;
import org.springframework.data.relational.core.mapping.Table;

import java.math.BigDecimal;
import java.util.UUID;

@Table("products")
public final class ProductRecord {

    @Id
    private final UUID identifier;
    private final BigDecimal price;
    private final String title;
    private final String description;

    public ProductRecord(UUID identifier, BigDecimal price, String title, String description) {
        this.identifier = identifier;
        this.price = price;
        this.title = title;
        this.description = description;
    }

    public UUID getIdentifier() {
        return identifier;
    }

    public BigDecimal getPrice() {
        return price;
    }

    public String getTitle() {
        return title;
    }

    public String getDescription() {
        return description;
    }
}
