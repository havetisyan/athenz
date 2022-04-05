//
// This file generated by rdl 1.5.2. Do not modify!
//

package com.yahoo.athenz.msd;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.List;
import com.yahoo.rdl.*;

//
// NetworkPolicyChangeImpactRequest - struct representing input details for
// evaluating network policies change impact on transport policies
//
@JsonIgnoreProperties(ignoreUnknown = true)
public class NetworkPolicyChangeImpactRequest {
    public List<IPBlock> from;
    public List<IPBlock> to;
    public List<NetworkPolicyPorts> ports;

    public NetworkPolicyChangeImpactRequest setFrom(List<IPBlock> from) {
        this.from = from;
        return this;
    }
    public List<IPBlock> getFrom() {
        return from;
    }
    public NetworkPolicyChangeImpactRequest setTo(List<IPBlock> to) {
        this.to = to;
        return this;
    }
    public List<IPBlock> getTo() {
        return to;
    }
    public NetworkPolicyChangeImpactRequest setPorts(List<NetworkPolicyPorts> ports) {
        this.ports = ports;
        return this;
    }
    public List<NetworkPolicyPorts> getPorts() {
        return ports;
    }

    @Override
    public boolean equals(Object another) {
        if (this != another) {
            if (another == null || another.getClass() != NetworkPolicyChangeImpactRequest.class) {
                return false;
            }
            NetworkPolicyChangeImpactRequest a = (NetworkPolicyChangeImpactRequest) another;
            if (from == null ? a.from != null : !from.equals(a.from)) {
                return false;
            }
            if (to == null ? a.to != null : !to.equals(a.to)) {
                return false;
            }
            if (ports == null ? a.ports != null : !ports.equals(a.ports)) {
                return false;
            }
        }
        return true;
    }
}
