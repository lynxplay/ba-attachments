package dev.lynxplay.ba.helidon;

import java.math.BigDecimal;
import java.util.UUID;

public final class ProductRecord {
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
