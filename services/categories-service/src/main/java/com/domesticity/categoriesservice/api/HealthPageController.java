package com.domesticity.categoriesservice.api;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.actuate.health.Health;
import org.springframework.boot.actuate.health.HealthIndicator;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.SingleColumnRowMapper;
import org.springframework.stereotype.Component;

@Component
public class HealthPageController implements HealthIndicator {

    @Autowired
    JdbcTemplate jdbcTemplate;

    @Override
    public Health health() {
        if (!checkService()) {
            return Health.down().build();
        }
        return Health.up().build();
    }

    private boolean checkService() {
        String queryStr = "select 1 from category";

        return jdbcTemplate.query(queryStr, new SingleColumnRowMapper<>()).size() > 0;
    }
}
