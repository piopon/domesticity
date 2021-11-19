package com.domesticity.categoriesservice.dao;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import com.domesticity.categoriesservice.model.Category;

import org.springframework.stereotype.Repository;

@Repository("memory")
public class InMemoryCategoryDao implements CategoryDao {

    private static final List<Category> MEMORY_DB = new ArrayList<>();

    @Override
    public int addCategory(String id, Category category) {
        MEMORY_DB.add(new Category(id, category.getName(), category.getColour(), category.getIcon()));
        return 1;
    }

    @Override
    public List<Category> getCategories() {
        return MEMORY_DB;
    }

    @Override
    public Optional<Category> getCategory(String id) {
        return MEMORY_DB.stream().filter(category -> category.getId().equals(id)).findFirst();
    }

    @Override
    public int deleteCategory(String id) {
        Optional<Category> foundCategory = getCategory(id);
        if (foundCategory.isEmpty()) {
            return 0;
        }
        MEMORY_DB.remove(foundCategory.get());
        return 1;
    }

    @Override
    public int updateCategory(String id, Category newCategory) {
        return getCategory(id).map(category -> {
            int toUpdateIndex = MEMORY_DB.indexOf(category);
            if (toUpdateIndex >= 0) {
                final Category updated = new Category(id, newCategory.getName(), newCategory.getColour(),
                        newCategory.getIcon());
                MEMORY_DB.set(toUpdateIndex, updated);
                return 1;
            }
            return 0;
        }).orElse(0);
    }
}
