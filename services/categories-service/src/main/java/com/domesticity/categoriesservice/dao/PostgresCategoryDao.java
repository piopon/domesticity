package com.domesticity.categoriesservice.dao;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;
import java.util.Optional;

import com.domesticity.categoriesservice.model.Category;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
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
            String sql = "INSERT INTO category (id, name, color, icon) VALUES (?, ?, ?, ?)";
            return jdbcTemplate.update(sql, id, category.getName(), category.getColor(), category.getIcon());
        }
        return 0;
    }

    @Override
    public List<Category> getAllCategories() {
        final String sql = "SELECT * FROM category";
        return jdbcTemplate.query(sql, new CategoryMapper());
    }

    @Override
    public List<Category> getFilteredCategories(String name, String color, String icon) {
        final String sql = "SELECT * FROM category WHERE name LIKE ? AND color LIKE ? AND icon LIKE ?";
        return jdbcTemplate.query(sql, new CategoryMapper(), adjustFilter(name), adjustFilter(color),
                adjustFilter(icon));
    }

    @Override
    public Optional<Category> getCategory(String id) {
        if (!isCategoryPresent(id)) {
            return Optional.empty();
        }
        final String sql = "SELECT * FROM category WHERE id = ?";
        Category category = jdbcTemplate.queryForObject(sql, new CategoryMapper(), id);
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
            String sql = "UPDATE category SET name = ?, icon = ?, color = ? WHERE id = ?";
            return jdbcTemplate.update(sql, newCategory.getName(), newCategory.getIcon(),
                    newCategory.getColor(), id);
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

    private class CategoryMapper implements RowMapper<Category> {

        @Override
        public Category mapRow(ResultSet resultSet, int row) throws SQLException {
            String id = resultSet.getString("id");
            String name = resultSet.getString("name");
            String color = resultSet.getString("color");
            String icon = resultSet.getString("icon");

            return new Category(id, name, color, icon);
        }
    }
}
