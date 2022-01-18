package com.domesticity.categoriesservice.api;

import com.domesticity.categoriesservice.utilities.DbUrlParser;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class HtmlPagesController {

    @Value("${repository.type}")
    private String repositoryType;
    @Value("${spring.datasource.url}")
    private String postgresUrl;

    @GetMapping("/")
    public String getHomePage(Model model) {
        DbUrlParser dbUrlParser = new DbUrlParser(postgresUrl);

        model.addAttribute("service_ver", "v1.0");
        model.addAttribute("spring_ver", "v2.5.8");
        model.addAttribute("build_date", "2022.01.16 23:54.00");
        model.addAttribute("commit_sha", "2c3507de1c78cffab0440d5596fca30c1be1dcc3");
        model.addAttribute("repo_type", repositoryType);
        model.addAttribute("sql_scheme", dbUrlParser.getScheme());
        model.addAttribute("sql_ip", dbUrlParser.getIP());
        model.addAttribute("sql_port", dbUrlParser.getPort());

        return "home";
    }
}
