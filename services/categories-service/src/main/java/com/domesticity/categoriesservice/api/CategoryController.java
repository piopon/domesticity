package com.domesticity.categoriesservice.api;

import java.util.List;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;

import com.domesticity.categoriesservice.model.Category;
import com.domesticity.categoriesservice.service.CategoryService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;

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

    @GetMapping(params = "name")
    public List<Category> getCategoriesByName(@RequestParam("name") String name) {
        return categoryService.getCategoriesByName(name);
    }

    @GetMapping(params = "color")
    public List<Category> getCategoriesByColor(@RequestParam("color") String color) {
        return categoryService.getCategoriesByColor(color);
    }

    @GetMapping(params = "icon")
    public List<Category> getCategoriesByIcon(@RequestParam("icon") String icon) {
        return categoryService.getCategoriesByIcon(icon);
    }

    @GetMapping(path = "{id}")
    public Category getCategory(@PathVariable("id") String id) {
        return categoryService.getCategory(id)
                .orElseThrow(() -> new ResponseStatusException(HttpStatus.NOT_FOUND, "entity not found"));
    }

    @DeleteMapping(path = "{id}")
    public void deleteCategory(@PathVariable("id") String id) {
        categoryService.deleteCategory(id);
    }
}
