package com.domesticity.categoriesservice.utilities;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public class UrlParserTest {

    private static final String TEST_INPUT = "jdbc:postgresql://localhost:5432/";
    private UrlParser testParser;

    @BeforeEach
    public void setup() {
        testParser = new UrlParser(TEST_INPUT);
    }

    @Test
    public void getSchemeReturnsCorrectString() {
        assertEquals("jdbc:postgresql", testParser.getScheme());
    }

    @Test
    public void getIPReturnsCorrectString() {
        assertEquals("localhost", testParser.getIP());
    }

    @Test
    public void getPortReturnsCorrectString() {
        assertEquals("5432", testParser.getPort());
    }
}
