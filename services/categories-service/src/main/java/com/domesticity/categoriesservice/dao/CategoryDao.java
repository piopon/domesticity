package com.domesticity.categoriesservice.dao;

import java.util.List;
import java.util.UUID;
import com.domesticity.categoriesservice.model.Category;

public interface CategoryDao {

    int addCategory(UUID id, Category category);

    List<Category> getCategories();

    default int addCategory(Category category) {
        final UUID id = UUID.randomUUID();
        return addCategory(id, category);
    }
}
