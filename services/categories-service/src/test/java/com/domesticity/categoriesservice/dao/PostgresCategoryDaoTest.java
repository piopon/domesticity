package com.domesticity.categoriesservice.dao;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.util.Optional;

import com.domesticity.categoriesservice.model.Category;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.datasource.embedded.EmbeddedDatabase;
import org.springframework.jdbc.datasource.embedded.EmbeddedDatabaseBuilder;
import org.springframework.jdbc.datasource.embedded.EmbeddedDatabaseType;

public class PostgresCategoryDaoTest {

    private PostgresCategoryDao testDao;

    @BeforeEach
    void setup() {
        EmbeddedDatabase db = new EmbeddedDatabaseBuilder().setType(EmbeddedDatabaseType.H2)
            .addScript("schema.sql")
            .addScript("test-data.sql")
            .build();
        JdbcTemplate jdbc = new JdbcTemplate(db);
        testDao = new PostgresCategoryDao(jdbc);
    }

    @Test
    void getCategoriesShouldRetrieveAllStoredCategories() {
        assertEquals(3, testDao.getCategories().size());
        assertEquals("1", testDao.getCategories().get(0).getId());
        assertEquals("green", testDao.getCategories().get(1).getName());
        assertEquals("#0000FF", testDao.getCategories().get(2).getColour());
    }

    @Test
    void getCategoryShouldRetrieveSelectedItemFromDb() {
        Optional<Category> actualItem = testDao.getCategory("1");

        assertTrue(actualItem.isPresent());
        assertEquals("red", actualItem.get().getName());
        assertEquals("#FF0000", actualItem.get().getColour());
        assertEquals("icon1", actualItem.get().getIcon());
    }
}
