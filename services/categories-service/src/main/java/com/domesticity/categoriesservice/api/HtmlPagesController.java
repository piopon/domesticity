package com.domesticity.categoriesservice.api;

import com.domesticity.categoriesservice.utilities.UrlParser;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.info.BuildProperties;
import org.springframework.core.SpringVersion;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class HtmlPagesController {

    @Autowired
    private BuildProperties buildProperties;
    @Value("${repository.type}")
    private String repositoryType;
    @Value("${spring.datasource.url}")
    private String postgresUrl;

    @GetMapping("/")
    public String getHomePage(Model model) {
        UrlParser urlParser = new UrlParser(postgresUrl);

        model.addAttribute("service_ver", "v" + buildProperties.getVersion());
        model.addAttribute("spring_ver", "v" + SpringVersion.getVersion());
        model.addAttribute("build_date", buildProperties.getTime());
        model.addAttribute("commit_sha", "2c3507de1c78cffab0440d5596fca30c1be1dcc3");
        model.addAttribute("repo_type", repositoryType);
        model.addAttribute("sql_scheme", urlParser.getScheme());
        model.addAttribute("sql_ip", urlParser.getIP());
        model.addAttribute("sql_port", urlParser.getPort());

        return "home";
    }
}
