package com.domesticity.categoriesservice;

import java.util.List;

import com.domesticity.categoriesservice.dao.CategoryDao;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import springfox.documentation.builders.PathSelectors;
import springfox.documentation.service.ApiInfo;
import springfox.documentation.service.Contact;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@Configuration
@EnableSwagger2
public class CategoriesServiceApplicationConfig {

    @Autowired
    private ApplicationContext context;

    @Bean
    public CategoryDao CategoryDaoRepository(@Value("${repository.type}") String qualifier) {
        return (CategoryDao) context.getBean(qualifier);
    }

    @Bean
    public Docket swaggerSettings() {
        return new Docket(DocumentationType.SWAGGER_2).select()
            .paths(PathSelectors.ant("/category/**/"))
            .build().apiInfo(getApiInfo());
    }

    private ApiInfo getApiInfo() {
        Contact contact = new Contact("Piotr Ponikowski", "", "piopon.github@gmail.com");
        String description = "This service is a part of domesticity web application. It's responsible" +
                             "for managing global categories for all types of notes.";
        return new ApiInfo("Categories service", description, "1.0", "", contact, "MIT", "", List.of());
    }
}
