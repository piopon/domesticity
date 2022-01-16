package com.domesticity.categoriesservice.config;

import java.util.List;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.ViewControllerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

import springfox.documentation.builders.PathSelectors;
import springfox.documentation.service.ApiInfo;
import springfox.documentation.service.Contact;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@Configuration
@EnableSwagger2
public class CategoriesServiceSwaggerConfig implements WebMvcConfigurer {

    @Bean
    public Docket swaggerSettings() {
        return new Docket(DocumentationType.SWAGGER_2).select()
                .paths(PathSelectors.ant("/category/**/"))
                .build().apiInfo(getApiInfo());
    }

    @Override
    public void addViewControllers(final ViewControllerRegistry registry) {
        registry.addRedirectViewController("/docs", "/docs/swagger-ui/");
    }

    private ApiInfo getApiInfo() {
        String title = "Domesticity Categories Service API";
        String description = "Documentation for Category Service API used by Domesticity web application.\n" +
                "This service is responsible for managing global categories settings for all types of notes/lists.";
        String version = "1.0";
        String docsUrl = "https://github.com/piopon/domesticity/blob/main/README.md";
        Contact contact = new Contact("Piotr Ponikowski", "", "piopon.github@gmail.com");
        String licenseName = "License: GNU General Public License v3.0 (GPL-3.0)";
        String licenseUrl = "https://github.com/piopon/domesticity/blob/main/LICENSE";

        return new ApiInfo(title, description, version, docsUrl, contact, licenseName, licenseUrl, List.of());
    }

}
