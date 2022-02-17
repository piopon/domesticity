package com.domesticity.categoriesservice.api;

import static org.mockito.Mockito.when;

import java.time.Instant;

import org.hamcrest.Matchers;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.info.BuildProperties;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultHandlers;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

@WebMvcTest(HomePageController.class)
public class HomePageControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private BuildProperties buildProperties;

    @Test
    public void handleIndexShouldReturnHomePage() throws Exception {
        final String testName = "artifact-name";
        final String testVer = "1.1.1";
        final Instant testTime = Instant.now();

        when(buildProperties.getArtifact()).thenReturn(testName);
        when(buildProperties.getVersion()).thenReturn(testVer);
        when(buildProperties.getTime()).thenReturn(testTime);

        mockMvc.perform(MockMvcRequestBuilders.get("/"))
                .andDo(MockMvcResultHandlers.print())
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.model().hasNoErrors())
                .andExpect(MockMvcResultMatchers.model().attribute("service_name", Matchers.equalTo(testName)))
                .andExpect(MockMvcResultMatchers.model().attribute("service_ver", Matchers.equalTo("v" + testVer)))
                .andExpect(MockMvcResultMatchers.model().attribute("build_date", Matchers.equalTo(testTime)))
                .andExpect(MockMvcResultMatchers.model().attributeExists("spring_ver", "commit_sha", "repo_type"))
                .andExpect(MockMvcResultMatchers.model().attributeExists("sql_scheme", "sql_ip", "sql_port"))
                .andExpect(MockMvcResultMatchers.content().string(Matchers.containsString("Commit SHA")));
    }
}
