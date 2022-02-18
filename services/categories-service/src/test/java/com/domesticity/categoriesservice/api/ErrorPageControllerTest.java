package com.domesticity.categoriesservice.api;

import static org.mockito.Mockito.when;

import org.hamcrest.Matchers;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.info.BuildProperties;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.boot.web.servlet.error.ErrorAttributes;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultHandlers;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

@WebMvcTest(ErrorPageController.class)
public class ErrorPageControllerTest {

    @Autowired
    private MockMvc mockMvc;
    @MockBean
    private ErrorAttributes errorAttributes;
    @MockBean
    private BuildProperties buildProperties;

    @Test
    public void handleErrorShouldReturnErrorPage() throws Exception {
        final String testName = "artifact-name";

        when(buildProperties.getArtifact()).thenReturn(testName);
        
        mockMvc.perform(MockMvcRequestBuilders.get("/error"))
                .andDo(MockMvcResultHandlers.print())
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.model().hasNoErrors())
                .andExpect(MockMvcResultMatchers.model().attribute("service_name", Matchers.equalTo(testName)))
                .andExpect(MockMvcResultMatchers.content().string(Matchers.containsString("URL path: ")));
    }
}
