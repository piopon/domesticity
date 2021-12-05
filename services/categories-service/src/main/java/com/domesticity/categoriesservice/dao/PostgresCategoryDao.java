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
        if (!isCategoryPresent(category.getName())) {
            String sql = "INSERT INTO category (id, name, colour, icon) VALUES (?, ?, ?, ?)";
            return jdbcTemplate.update(sql, id, category.getName(), category.getColour(), category.getIcon());
        }
        return 0;
    }

    @Override
    public List<Category> getCategories() {
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
    public Optional<Category> getCategory(String id) {
        final String sql = "SELECT * FROM category cat WHERE cat.id = ?";
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
        if (isCategoryPresent(newCategory.getName())) {
            String sql = "UPDATE category SET name = ?, icon = ?, colour = ? WHERE id = ?";
            return jdbcTemplate.update(sql, newCategory.getName(), newCategory.getIcon(), newCategory.getColour(), id);
        }
        return 0;
    }

    private boolean isCategoryPresent(String name) {
        String sql = "SELECT EXISTS (SELECT 1 FROM category WHERE name = ?)";
        return jdbcTemplate.queryForObject(sql, (resultSet, i) -> resultSet.getBoolean(1), name);
    }
}
