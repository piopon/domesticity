package com.domesticity.categoriesservice.utilities;

public class UrlParser {

    private static final String SCHEME_DIVIDER = "://";
    private static final String PORT_DIVIDER = ":";
    private static final String ENDPOINT_DIVIDER = "/";

    private final String scheme;
    private final String ipAddress;
    private final String portNumber;

    public UrlParser(final String input) {
        final int schemeDividerIndex = input.indexOf(SCHEME_DIVIDER);
        final int portDividerIndex = input.indexOf(PORT_DIVIDER, schemeDividerIndex + 1);
        final int endpointDividerIndex = input.indexOf(ENDPOINT_DIVIDER, portDividerIndex + 1);

        scheme = input.substring(0, schemeDividerIndex);
        ipAddress = input.substring(schemeDividerIndex + SCHEME_DIVIDER.length(), portDividerIndex);
        portNumber = input.substring(portDividerIndex + PORT_DIVIDER.length(), endpointDividerIndex);
    }

    public String getScheme() {
        return scheme;
    }

    public String getIP() {
        return ipAddress;
    }

    public String getPort() {
        return portNumber;
    }
}
