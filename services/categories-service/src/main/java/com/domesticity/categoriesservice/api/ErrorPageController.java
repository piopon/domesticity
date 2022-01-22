package com.domesticity.categoriesservice.api;

import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.info.BuildProperties;
import org.springframework.boot.web.error.ErrorAttributeOptions;
import org.springframework.boot.web.servlet.error.ErrorAttributes;
import org.springframework.boot.web.servlet.error.ErrorController;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.context.request.ServletWebRequest;

@Controller
public class ErrorPageController implements ErrorController {

    @Autowired
    private ErrorAttributes errorAttributes;
    @Autowired
    private BuildProperties buildProperties;

    @RequestMapping("/error")
    public String handleError(Model model, HttpServletRequest webRequest) {
        model.addAttribute("service_name", buildProperties.getArtifact());
        ServletWebRequest servletWebRequest = new ServletWebRequest(webRequest);
        Map<String, Object> attributes = errorAttributes.getErrorAttributes(servletWebRequest, ErrorAttributeOptions.defaults());
        attributes.forEach((attribute, value) -> {
            model.addAttribute(attribute, value);
        });

        return "error";
    }
}
