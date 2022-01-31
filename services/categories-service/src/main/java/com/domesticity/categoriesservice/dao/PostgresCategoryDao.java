package com.domesticity.categoriesservice.dao;

import java.util.List;
import java.util.Optional;

import com.domesticity.categoriesservice.model.Category;

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
        if (!isCategoryPresent(id)) {
            String sql = "INSERT INTO category (id, name, colour, icon) VALUES (?, ?, ?, ?)";
            return jdbcTemplate.update(sql, id, category.getName(), category.getColour(), category.getIcon());
        }
        return 0;
    }

    @Override
    public List<Category> getAllCategories() {
        final String sql = "SELECT * FROM category";
        return jdbcTemplate.query(sql, (results, i) -> {
            String id = results.getString("id");
            String name = results.getString("name");
            String colour = results.getString("colour");
            String icon = results.getString("icon");
            return new Category(id, name, colour, icon);
        });
    }

    @Override
    public List<Category> getFilteredCategories(String name, String color, String icon) {
        final String sql = "SELECT * FROM category WHERE name LIKE ? AND colour LIKE ? AND icon LIKE ?";
        return jdbcTemplate.query(sql, (results, i) -> {
            String idStr = results.getString("id");
            String nameStr = results.getString("name");
            String colorStr = results.getString("colour");
            String iconStr = results.getString("icon");
            return new Category(idStr, nameStr, colorStr, iconStr);
        }, adjustFilter(name), adjustFilter(color), adjustFilter(icon));
    }

    @Override
    public Optional<Category> getCategory(String id) {
        if (!isCategoryPresent(id)) {
            return Optional.empty();
        }
        final String sql = "SELECT * FROM category WHERE id = ?";
        Category category = jdbcTemplate.queryForObject(sql, (results, i) -> {
            String name = results.getString("name");
            String colour = results.getString("colour");
            String icon = results.getString("icon");
            return new Category(id, name, colour, icon);
        }, id);
        return Optional.ofNullable(category);
    }

    @Override
    public int deleteCategory(String id) {
        final String sql = "DELETE FROM category WHERE id = ?";
        return jdbcTemplate.update(sql, id);
    }

    @Override
    public int updateCategory(String id, Category newCategory) {
        if (isCategoryPresent(id)) {
            String sql = "UPDATE category SET name = ?, icon = ?, colour = ? WHERE id = ?";
            return jdbcTemplate.update(sql, newCategory.getName(), newCategory.getIcon(), newCategory.getColour(), id);
        }
        return 0;
    }

    private boolean isCategoryPresent(String id) {
        String sql = "SELECT EXISTS (SELECT 1 FROM category WHERE id = ?)";
        return jdbcTemplate.queryForObject(sql, (results, i) -> results.getBoolean(1), id);
    }

    private String adjustFilter(final String input) {
        return (input == null) ? "%" : input;
    }
}
