package com.domesticity.categoriesservice.dao;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

import com.domesticity.categoriesservice.model.Category;

import org.springframework.stereotype.Repository;

@Repository("memory")
public class InMemoryCategoryDao implements CategoryDao {

    private static final List<Category> MEMORY_DB = new ArrayList<>();

    @Override
    public int addCategory(UUID id, Category category) {
        MEMORY_DB.add(new Category(id, category.getColour(), category.getIcon()));
        return 1;
    }
}
