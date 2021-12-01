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
        addColour(category.getColour());
        if (!isCategoryPresent(category.getName())) {
            String sql = "INSERT INTO category (id, name, colour, icon) VALUES (?, ?, ?, ?)";
            return jdbcTemplate.update(sql, id, category.getName(), category.getColour().getName(), category.getIcon());
        }
        return 0;
    }

    @Override
    public List<Category> getCategories() {
        final String sql = "SELECT * FROM category cat, colour col WHERE cat.colour = col.name";
        return jdbcTemplate.query(sql, (results, i) -> {
            String id = results.getString("id");
            String name = results.getString("name");
            String icon = results.getString("icon");
            final Colour colour = new Colour(results.getString("colour"),
                                             results.getString("red"),
                                             results.getString("green"),
                                             results.getString("blue"),
                                             results.getString("alpha"));
            return new Category(id, name, colour, icon);
        });
    }

    @Override
    public Optional<Category> getCategory(String id) {
        final String sql = "SELECT * FROM category cat, colour col WHERE cat.colour = col.name AND cat.id = ?";
        Category category = jdbcTemplate.queryForObject(sql, (results, i) -> {
            String name = results.getString("name");
            String icon = results.getString("icon");
            final Colour colour = new Colour(results.getString("colour"),
                                             results.getString("red"),
                                             results.getString("green"),
                                             results.getString("blue"),
                                             results.getString("alpha"));
            return new Category(id, name, colour, icon);
        }, id);
        return Optional.ofNullable(category);
    }

    @Override
    public int deleteCategory(String id) {
        final String sql = "DELETE FROM student WHERE student_id = ?";
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

    private int addColour(Colour colour) {
        if (!isColourPresent(colour.getName())) {
            String sql = "INSERT INTO colour (name, red, green, blue, alpha) VALUES (?, ?, ?, ?, ?)";
            return jdbcTemplate.update(sql, colour.getName(), colour.getRed(), colour.getGreen(), colour.getBlue(),
                    colour.getAlpha());
        }
        return 0;
    }

    private boolean isColourPresent(String name) {
        String sql = "SELECT EXISTS (SELECT 1 FROM colour WHERE name = ?)";
        return jdbcTemplate.queryForObject(sql, (resultSet, i) -> resultSet.getBoolean(1), name);
    }

    private boolean isCategoryPresent(String name) {
        String sql = "SELECT EXISTS (SELECT 1 FROM category WHERE name = ?)";
        return jdbcTemplate.queryForObject(sql, (resultSet, i) -> resultSet.getBoolean(1), name);
    }
}
