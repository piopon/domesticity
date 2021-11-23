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
        final String sql = "SELECT * FROM category cat, colour col WHERE cat.colour = col.name";
        return jdbcTemplate.query(sql, (results, i) -> {
            String id = results.getString("id");
            String name = results.getString("name");
            String icon = results.getString("icon");
            final Colour colour = new Colour(results.getString("red"),
                                             results.getString("green"),
                                             results.getString("blue"),
                                             results.getString("alpha"));
            return new Category(id, name, colour, icon);
        });
    }

    @Override
    public Optional<Category> getCategory(String id) {
        final String sql = "SELECT * FROM category cat, colour col WHERE cat.id = ? AND cat.colour = col.name";
        Category category = jdbcTemplate.queryForObject(sql, new Object[]{id}, (results, i) -> {
            String name = results.getString("name");
            String icon = results.getString("icon");
            final Colour colour = new Colour(results.getString("red"),
                                             results.getString("green"),
                                             results.getString("blue"),
                                             results.getString("alpha"));
            return new Category(id, name, colour, icon);
        });
        return Optional.ofNullable(category);
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
