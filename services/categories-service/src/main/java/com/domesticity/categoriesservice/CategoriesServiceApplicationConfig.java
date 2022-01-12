package com.domesticity.categoriesservice;

import com.domesticity.categoriesservice.dao.CategoryDao;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

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
}
