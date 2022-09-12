// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.grafana.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeamState extends com.pulumi.resources.ResourceArgs {

    public static final TeamState Empty = new TeamState();

    /**
     * An email address for the team.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return An email address for the team.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * A set of email addresses corresponding to users who should be given membership
     * to the team. Note: users specified here must already exist in Grafana.
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<String>> members;

    /**
     * @return A set of email addresses corresponding to users who should be given membership
     * to the team. Note: users specified here must already exist in Grafana.
     * 
     */
    public Optional<Output<List<String>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The display name for the Grafana team created.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The display name for the Grafana team created.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The team id assigned to this team by Grafana.
     * 
     */
    @Import(name="teamId")
    private @Nullable Output<Integer> teamId;

    /**
     * @return The team id assigned to this team by Grafana.
     * 
     */
    public Optional<Output<Integer>> teamId() {
        return Optional.ofNullable(this.teamId);
    }

    private TeamState() {}

    private TeamState(TeamState $) {
        this.email = $.email;
        this.members = $.members;
        this.name = $.name;
        this.teamId = $.teamId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamState $;

        public Builder() {
            $ = new TeamState();
        }

        public Builder(TeamState defaults) {
            $ = new TeamState(Objects.requireNonNull(defaults));
        }

        /**
         * @param email An email address for the team.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email An email address for the team.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param members A set of email addresses corresponding to users who should be given membership
         * to the team. Note: users specified here must already exist in Grafana.
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<String>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members A set of email addresses corresponding to users who should be given membership
         * to the team. Note: users specified here must already exist in Grafana.
         * 
         * @return builder
         * 
         */
        public Builder members(List<String> members) {
            return members(Output.of(members));
        }

        /**
         * @param members A set of email addresses corresponding to users who should be given membership
         * to the team. Note: users specified here must already exist in Grafana.
         * 
         * @return builder
         * 
         */
        public Builder members(String... members) {
            return members(List.of(members));
        }

        /**
         * @param name The display name for the Grafana team created.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The display name for the Grafana team created.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param teamId The team id assigned to this team by Grafana.
         * 
         * @return builder
         * 
         */
        public Builder teamId(@Nullable Output<Integer> teamId) {
            $.teamId = teamId;
            return this;
        }

        /**
         * @param teamId The team id assigned to this team by Grafana.
         * 
         * @return builder
         * 
         */
        public Builder teamId(Integer teamId) {
            return teamId(Output.of(teamId));
        }

        public TeamState build() {
            return $;
        }
    }

}
