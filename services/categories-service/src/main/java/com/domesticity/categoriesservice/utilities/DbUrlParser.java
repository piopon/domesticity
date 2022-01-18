package com.domesticity.categoriesservice.utilities;

public class DbUrlParser {

    private static final String SCHEME_DIVIDER = "://";
    private static final String PORT_DIVIDER = ":";
    private static final String DB_DIVIDER = "/";

    private String scheme;
    private String ipAddress;
    private String portNumber;

    public DbUrlParser(String input) {
        int schemeDividerIndex = input.indexOf(SCHEME_DIVIDER);
        System.out.println("scheme index = " + schemeDividerIndex);
        int portDividerIndex = input.indexOf(PORT_DIVIDER, schemeDividerIndex + 1);
        System.out.println("port index = " + portDividerIndex);
        int dbNameDividerIndex = input.indexOf(DB_DIVIDER, portDividerIndex + 1);
        System.out.println("db index = " + dbNameDividerIndex);

        scheme = input.substring(0, schemeDividerIndex);
        System.out.println("scheme = " + scheme);
        ipAddress = input.substring(schemeDividerIndex + SCHEME_DIVIDER.length(), portDividerIndex);
        System.out.println("ipAddress = " + ipAddress);
        portNumber = input.substring(portDividerIndex + PORT_DIVIDER.length(), dbNameDividerIndex);
        System.out.println("port = " + portNumber);
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
