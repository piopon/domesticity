package com.domesticity.categoriesservice.api;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

import static org.mockito.Mockito.when;

import java.util.List;

import com.domesticity.categoriesservice.model.Category;
import com.domesticity.categoriesservice.service.CategoryService;

@SpringBootTest
@AutoConfigureMockMvc
public class CategoryControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private CategoryService mockService;

    @Test
    void receivingAllCategoriesWithCorrectUrlReturnsOkStatus() throws Exception {
        when(mockService.getCategories()).thenReturn(List.of(
                new Category("1", "kategoria1", "#0000FF", "ikona1")));

        mockMvc.perform(MockMvcRequestBuilders.get("/category"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.jsonPath("$.size()").value(1));
    }

    @Test
    void deletingSingleCategoryWithCorrectUrlReturnsOkStatus() throws Exception {
        mockMvc.perform(MockMvcRequestBuilders.delete("/category/123")
                .contentType("application/json"))
                .andExpect(MockMvcResultMatchers.status().isOk());
    }
}
