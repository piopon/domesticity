package com.domesticity.categoriesservice.dao;

import java.util.List;
import java.util.Optional;

import com.domesticity.categoriesservice.model.Category;
import com.domesticity.categoriesservice.model.Colour;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

@Repository("postgres")
public class PostgresCategoryDao implements CategoryDao {

    private final static Colour TEMP_COLOUR = new Colour("200","200","200","1");
    private final JdbcTemplate jdbcTemplate;

    @Autowired
    public PostgresCategoryDao(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @Override
    public int addCategory(String id, Category category) {
        return 0;
    }

    @Override
    public List<Category> getCategories() {
        final String sql = "SELECT id, name, colour, icon FROM category";
        return jdbcTemplate.query(sql, (results, i) -> {
            String id = results.getString("id");
            String name = results.getString("name");
            String icon = results.getString("icon");
            return new Category(id, name, TEMP_COLOUR, icon);
        });
    }

    @Override
    public Optional<Category> getCategory(String id) {
        return null;
    }

    @Override
    public int deleteCategory(String id) {
        return 0;
    }

    @Override
    public int updateCategory(String id, Category newCategory) {
        return 0;
    }
}
