package com.domesticity.categoriesservice.api;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class HtmlPagesController {

    @GetMapping("/")
    String getHomePage(Model model) {
        model.addAttribute("service_ver", "v1.0");
        model.addAttribute("spring_ver", "v2.5.8");
        model.addAttribute("build_date", "2022.01.16 23:54.00");
        model.addAttribute("commit_sha", "2c3507de1c78cffab0440d5596fca30c1be1dcc3");
        model.addAttribute("repo_type", "postgres");
        model.addAttribute("sql_scheme", "jdbc:postgresql");
        model.addAttribute("sql_ip", "127.0.0.1");
        model.addAttribute("sql_port", "5432");

        return "home";
    }
}
