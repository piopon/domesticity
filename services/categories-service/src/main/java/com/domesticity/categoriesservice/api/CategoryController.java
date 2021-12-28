package com.domesticity.categoriesservice.api;

import java.util.List;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;

import com.domesticity.categoriesservice.model.Category;
import com.domesticity.categoriesservice.service.CategoryService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("category")
public class CategoryController {

    private final CategoryService categoryService;

    @Autowired
    public CategoryController(CategoryService categoryService) {
        this.categoryService = categoryService;
    }

    @PostMapping
    public void addCategory(@Valid @NotNull @RequestBody Category category) {
        categoryService.addCategory(category);
    }

    @PutMapping(path = "{id}")
    public void updateCategory(@PathVariable("id") String id,
                               @Valid @NotNull @RequestBody Category category) {
        categoryService.updateCategory(id, category);
    }

    @GetMapping
    public List<Category> getCategories() {
        return categoryService.getCategories();
    }

    @GetMapping(path = "{id}")
    public Category getCategory(@PathVariable("id") String id) {
        return categoryService.getCategory(id).orElse(null);
    }

    @DeleteMapping(path = "{id}")
    public void deleteCategory(@PathVariable("id") String id) {
        categoryService.deleteCategory(id);
    }
}
