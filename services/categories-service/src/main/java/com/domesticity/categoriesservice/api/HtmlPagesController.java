package com.domesticity.categoriesservice.api;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class HtmlPagesController {

    @GetMapping("/")
    String getHomePage(Model model) {
        model.addAttribute("service-version", "v1.0");
        model.addAttribute("spring-version", "v2.5.8");
        model.addAttribute("build-time", "16.01.2022 23:54.00");

        return "home";
    }
}
