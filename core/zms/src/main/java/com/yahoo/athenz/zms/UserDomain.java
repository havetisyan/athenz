//
// This file generated by rdl 1.4.10. Do not modify!
//

package com.yahoo.athenz.zms;
import com.yahoo.rdl.*;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

//
// UserDomain - A UserDomain is the user's own top level domain in user - e.g.
// user.hga
//
@JsonSerialize(include = JsonSerialize.Inclusion.NON_DEFAULT)
public class UserDomain {
    @RdlOptional
    public String description;
    @RdlOptional
    public String org;
    @RdlOptional
    public Boolean enabled;
    @RdlOptional
    public Boolean auditEnabled;
    @RdlOptional
    public String account;
    @RdlOptional
    public Integer ypmId;
    public String name;
    @RdlOptional
    public DomainTemplateList templates;

    public UserDomain setDescription(String description) {
        this.description = description;
        return this;
    }
    public String getDescription() {
        return description;
    }
    public UserDomain setOrg(String org) {
        this.org = org;
        return this;
    }
    public String getOrg() {
        return org;
    }
    public UserDomain setEnabled(Boolean enabled) {
        this.enabled = enabled;
        return this;
    }
    public Boolean getEnabled() {
        return enabled;
    }
    public UserDomain setAuditEnabled(Boolean auditEnabled) {
        this.auditEnabled = auditEnabled;
        return this;
    }
    public Boolean getAuditEnabled() {
        return auditEnabled;
    }
    public UserDomain setAccount(String account) {
        this.account = account;
        return this;
    }
    public String getAccount() {
        return account;
    }
    public UserDomain setYpmId(Integer ypmId) {
        this.ypmId = ypmId;
        return this;
    }
    public Integer getYpmId() {
        return ypmId;
    }
    public UserDomain setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return name;
    }
    public UserDomain setTemplates(DomainTemplateList templates) {
        this.templates = templates;
        return this;
    }
    public DomainTemplateList getTemplates() {
        return templates;
    }

    @Override
    public boolean equals(Object another) {
        if (this != another) {
            if (another == null || another.getClass() != UserDomain.class) {
                return false;
            }
            UserDomain a = (UserDomain) another;
            if (description == null ? a.description != null : !description.equals(a.description)) {
                return false;
            }
            if (org == null ? a.org != null : !org.equals(a.org)) {
                return false;
            }
            if (enabled == null ? a.enabled != null : !enabled.equals(a.enabled)) {
                return false;
            }
            if (auditEnabled == null ? a.auditEnabled != null : !auditEnabled.equals(a.auditEnabled)) {
                return false;
            }
            if (account == null ? a.account != null : !account.equals(a.account)) {
                return false;
            }
            if (ypmId == null ? a.ypmId != null : !ypmId.equals(a.ypmId)) {
                return false;
            }
            if (name == null ? a.name != null : !name.equals(a.name)) {
                return false;
            }
            if (templates == null ? a.templates != null : !templates.equals(a.templates)) {
                return false;
            }
        }
        return true;
    }
}
