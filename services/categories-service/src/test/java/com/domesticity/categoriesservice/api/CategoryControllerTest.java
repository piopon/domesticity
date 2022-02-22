package com.domesticity.categoriesservice.api;

import java.util.List;
import java.util.Optional;

import com.domesticity.categoriesservice.model.Category;
import com.domesticity.categoriesservice.service.CategoryService;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@WebMvcTest(CategoryController.class)
public class CategoryControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private CategoryService mockService;

    @Test
    public void receivingAllCategoriesWithCorrectUrlReturnsOkStatus() throws Exception {
        when(mockService.getAllCategories()).thenReturn(List.of(
                new Category("1", "kategoria1", "#0000FF", "ikona1"),
                new Category("2", "kategoria2", "#00FF00", "ikona2"),
                new Category("3", "kategoria3", "#FF0000", "ikona3")));

        mockMvc.perform(MockMvcRequestBuilders.get("/category"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.jsonPath("$.size()").value(3))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].id").value("1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[1].name").value("kategoria2"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[2].color").value("#FF0000"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].icon").value("ikona1"));
    }

    @Test
    public void receivingFilteredCategoriesWithAllFiltersInUrlReturnsOkStatus() throws Exception {
        when(mockService.getFilteredCategories("kategoria1", "blue", "ikona1")).thenReturn(List.of(
                new Category("1", "kategoria1", "blue", "ikona1")));

        mockMvc.perform(MockMvcRequestBuilders.get("/category?name=kategoria1&color=blue&icon=ikona1"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.jsonPath("$.size()").value(1))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].id").value("1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].name").value("kategoria1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].color").value("blue"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].icon").value("ikona1"));
    }

    @Test
    public void receivingFilteredCategoriesWithTwoFiltersInUrlReturnsOkStatus() throws Exception {
        when(mockService.getFilteredCategories(null, "blue", "ikona1")).thenReturn(List.of(
                new Category("1", "kategoria1", "blue", "ikona1")));

        mockMvc.perform(MockMvcRequestBuilders.get("/category?color=blue&icon=ikona1"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.jsonPath("$.size()").value(1))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].id").value("1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].name").value("kategoria1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].color").value("blue"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].icon").value("ikona1"));
    }

    @Test
    public void receivingFilteredCategoriesWithNameFilterInUrlReturnsOkStatus() throws Exception {
        when(mockService.getFilteredCategories("kategoria1", null, null)).thenReturn(List.of(
                new Category("1", "kategoria1", "blue", "ikona1")));

        mockMvc.perform(MockMvcRequestBuilders.get("/category?name=kategoria1"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.jsonPath("$.size()").value(1))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].id").value("1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].name").value("kategoria1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].color").value("blue"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].icon").value("ikona1"));
    }

    @Test
    public void receivingFilteredCategoriesWithColorFilterInUrlReturnsOkStatus() throws Exception {
        when(mockService.getFilteredCategories(null, "blue", null)).thenReturn(List.of(
                new Category("1", "kategoria1", "blue", "ikona1")));

        mockMvc.perform(MockMvcRequestBuilders.get("/category?color=blue"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.jsonPath("$.size()").value(1))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].id").value("1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].name").value("kategoria1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].color").value("blue"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].icon").value("ikona1"));
    }

    @Test
    public void receivingFilteredCategoriesWithIconFilterInUrlReturnsOkStatus() throws Exception {
        when(mockService.getFilteredCategories(null, null, "ikona1")).thenReturn(List.of(
                new Category("1", "kategoria1", "blue", "ikona1")));

        mockMvc.perform(MockMvcRequestBuilders.get("/category?icon=ikona1"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.jsonPath("$.size()").value(1))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].id").value("1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].name").value("kategoria1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].color").value("blue"))
                .andExpect(MockMvcResultMatchers.jsonPath("$[0].icon").value("ikona1"));
    }

    @Test
    public void receivingSingleCategoryWithCorrectUrlReturnsOkStatus() throws Exception {
        when(mockService.getCategory("123")).thenReturn(Optional.of(
                new Category("123", "cat1", "#123456", "bball")));

        mockMvc.perform(MockMvcRequestBuilders.get("/category/123"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.jsonPath("$.id").value("123"))
                .andExpect(MockMvcResultMatchers.jsonPath("$.name").value("cat1"))
                .andExpect(MockMvcResultMatchers.jsonPath("$.color").value("#123456"))
                .andExpect(MockMvcResultMatchers.jsonPath("$.icon").value("bball"));
    }

    @Test
    public void receivingSingleCategoryWithIncorrectUrlReturnsError() throws Exception {
        when(mockService.getCategory("123")).thenReturn(Optional.of(
                new Category("123", "cat1", "#123456", "bball")));

        mockMvc.perform(MockMvcRequestBuilders.get("/category/000"))
                .andExpect(MockMvcResultMatchers.status().isNotFound());
    }

    @Test
    public void addingSingleCategoryWithCorrectUrlReturnsOkStatus() throws Exception {
        mockMvc.perform(MockMvcRequestBuilders.post("/category")
                .contentType(MediaType.APPLICATION_JSON)
                .content("{\"name\": \"kat2\", \"color\":\"red\", \"icon\":\"zzz\"}"))
                .andExpect(MockMvcResultMatchers.status().isOk());
        verify(mockService).addCategory(any(Category.class));
    }

    @Test
    public void updatingSingleCategoryWithCorrectUrlReturnsOkStatus() throws Exception {
        mockMvc.perform(MockMvcRequestBuilders.put("/category/123")
                .contentType(MediaType.APPLICATION_JSON)
                .content("{\"name\": \"kat2\", \"color\":\"red\", \"icon\":\"zzz\"}"))
                .andExpect(MockMvcResultMatchers.status().isOk());
        verify(mockService).updateCategory(eq("123"), any(Category.class));
    }

    @Test
    public void deletingSingleCategoryWithCorrectUrlReturnsOkStatus() throws Exception {
        when(mockService.deleteCategory("123")).thenReturn(1);

        mockMvc.perform(MockMvcRequestBuilders.delete("/category/123")
                .contentType(MediaType.APPLICATION_JSON))
                .andExpect(MockMvcResultMatchers.status().isOk());
        verify(mockService).deleteCategory(eq("123"));
    }
}
