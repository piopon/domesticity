package com.domesticity.categoriesservice.api;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.web.servlet.MockMvc;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.delete;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest
@AutoConfigureMockMvc
public class CategoryControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @Test
    void receivingAllCategoriesWithCorrectUrlReturnsOkStatus() throws Exception {
        mockMvc.perform(get("/category")
                .contentType("application/json"))
                .andExpect(status().isOk());
    }

    @Test
    void deletingSingleCategoryWithCorrectUrlReturnsOkStatus() throws Exception {
        mockMvc.perform(delete("/category/123")
                .contentType("application/json"))
                .andExpect(status().isOk());
    }
}
