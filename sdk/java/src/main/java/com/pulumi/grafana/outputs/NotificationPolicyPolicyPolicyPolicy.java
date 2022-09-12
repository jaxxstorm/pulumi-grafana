// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.grafana.outputs.NotificationPolicyPolicyPolicyPolicyMatcher;
import com.pulumi.grafana.outputs.NotificationPolicyPolicyPolicyPolicyPolicy;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class NotificationPolicyPolicyPolicyPolicy {
    /**
     * @return The contact point to route notifications that match this rule to.
     * 
     */
    private String contactPoint;
    /**
     * @return Whether to continue matching subsequent rules if an alert matches the current rule. Otherwise, the rule will be &#39;consumed&#39; by the first policy to match it.
     * 
     */
    private @Nullable Boolean continue_;
    /**
     * @return A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping.
     * 
     */
    private List<String> groupBies;
    /**
     * @return Minimum time interval between two notifications for the same group. Default is 5 minutes.
     * 
     */
    private @Nullable String groupInterval;
    /**
     * @return Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
     * 
     */
    private @Nullable String groupWait;
    /**
     * @return Describes which labels this rule should match. When multiple matchers are supplied, an alert must match ALL matchers to be accepted by this policy. When no matchers are supplied, the rule will match all alert instances.
     * 
     */
    private @Nullable List<NotificationPolicyPolicyPolicyPolicyMatcher> matchers;
    /**
     * @return A list of mute timing names to apply to alerts that match this policy.
     * 
     */
    private @Nullable List<String> muteTimings;
    /**
     * @return Routing rules for specific label sets.
     * 
     */
    private @Nullable List<NotificationPolicyPolicyPolicyPolicyPolicy> policies;
    /**
     * @return Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
     * 
     */
    private @Nullable String repeatInterval;

    private NotificationPolicyPolicyPolicyPolicy() {}
    /**
     * @return The contact point to route notifications that match this rule to.
     * 
     */
    public String contactPoint() {
        return this.contactPoint;
    }
    /**
     * @return Whether to continue matching subsequent rules if an alert matches the current rule. Otherwise, the rule will be &#39;consumed&#39; by the first policy to match it.
     * 
     */
    public Optional<Boolean> continue_() {
        return Optional.ofNullable(this.continue_);
    }
    /**
     * @return A list of alert labels to group alerts into notifications by. Use the special label `...` to group alerts by all labels, effectively disabling grouping.
     * 
     */
    public List<String> groupBies() {
        return this.groupBies;
    }
    /**
     * @return Minimum time interval between two notifications for the same group. Default is 5 minutes.
     * 
     */
    public Optional<String> groupInterval() {
        return Optional.ofNullable(this.groupInterval);
    }
    /**
     * @return Time to wait to buffer alerts of the same group before sending a notification. Default is 30 seconds.
     * 
     */
    public Optional<String> groupWait() {
        return Optional.ofNullable(this.groupWait);
    }
    /**
     * @return Describes which labels this rule should match. When multiple matchers are supplied, an alert must match ALL matchers to be accepted by this policy. When no matchers are supplied, the rule will match all alert instances.
     * 
     */
    public List<NotificationPolicyPolicyPolicyPolicyMatcher> matchers() {
        return this.matchers == null ? List.of() : this.matchers;
    }
    /**
     * @return A list of mute timing names to apply to alerts that match this policy.
     * 
     */
    public List<String> muteTimings() {
        return this.muteTimings == null ? List.of() : this.muteTimings;
    }
    /**
     * @return Routing rules for specific label sets.
     * 
     */
    public List<NotificationPolicyPolicyPolicyPolicyPolicy> policies() {
        return this.policies == null ? List.of() : this.policies;
    }
    /**
     * @return Minimum time interval for re-sending a notification if an alert is still firing. Default is 4 hours.
     * 
     */
    public Optional<String> repeatInterval() {
        return Optional.ofNullable(this.repeatInterval);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(NotificationPolicyPolicyPolicyPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String contactPoint;
        private @Nullable Boolean continue_;
        private List<String> groupBies;
        private @Nullable String groupInterval;
        private @Nullable String groupWait;
        private @Nullable List<NotificationPolicyPolicyPolicyPolicyMatcher> matchers;
        private @Nullable List<String> muteTimings;
        private @Nullable List<NotificationPolicyPolicyPolicyPolicyPolicy> policies;
        private @Nullable String repeatInterval;
        public Builder() {}
        public Builder(NotificationPolicyPolicyPolicyPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.contactPoint = defaults.contactPoint;
    	      this.continue_ = defaults.continue_;
    	      this.groupBies = defaults.groupBies;
    	      this.groupInterval = defaults.groupInterval;
    	      this.groupWait = defaults.groupWait;
    	      this.matchers = defaults.matchers;
    	      this.muteTimings = defaults.muteTimings;
    	      this.policies = defaults.policies;
    	      this.repeatInterval = defaults.repeatInterval;
        }

        @CustomType.Setter
        public Builder contactPoint(String contactPoint) {
            this.contactPoint = Objects.requireNonNull(contactPoint);
            return this;
        }
        @CustomType.Setter("continue")
        public Builder continue_(@Nullable Boolean continue_) {
            this.continue_ = continue_;
            return this;
        }
        @CustomType.Setter
        public Builder groupBies(List<String> groupBies) {
            this.groupBies = Objects.requireNonNull(groupBies);
            return this;
        }
        public Builder groupBies(String... groupBies) {
            return groupBies(List.of(groupBies));
        }
        @CustomType.Setter
        public Builder groupInterval(@Nullable String groupInterval) {
            this.groupInterval = groupInterval;
            return this;
        }
        @CustomType.Setter
        public Builder groupWait(@Nullable String groupWait) {
            this.groupWait = groupWait;
            return this;
        }
        @CustomType.Setter
        public Builder matchers(@Nullable List<NotificationPolicyPolicyPolicyPolicyMatcher> matchers) {
            this.matchers = matchers;
            return this;
        }
        public Builder matchers(NotificationPolicyPolicyPolicyPolicyMatcher... matchers) {
            return matchers(List.of(matchers));
        }
        @CustomType.Setter
        public Builder muteTimings(@Nullable List<String> muteTimings) {
            this.muteTimings = muteTimings;
            return this;
        }
        public Builder muteTimings(String... muteTimings) {
            return muteTimings(List.of(muteTimings));
        }
        @CustomType.Setter
        public Builder policies(@Nullable List<NotificationPolicyPolicyPolicyPolicyPolicy> policies) {
            this.policies = policies;
            return this;
        }
        public Builder policies(NotificationPolicyPolicyPolicyPolicyPolicy... policies) {
            return policies(List.of(policies));
        }
        @CustomType.Setter
        public Builder repeatInterval(@Nullable String repeatInterval) {
            this.repeatInterval = repeatInterval;
            return this;
        }
        public NotificationPolicyPolicyPolicyPolicy build() {
            final var o = new NotificationPolicyPolicyPolicyPolicy();
            o.contactPoint = contactPoint;
            o.continue_ = continue_;
            o.groupBies = groupBies;
            o.groupInterval = groupInterval;
            o.groupWait = groupWait;
            o.matchers = matchers;
            o.muteTimings = muteTimings;
            o.policies = policies;
            o.repeatInterval = repeatInterval;
            return o;
        }
    }
}
