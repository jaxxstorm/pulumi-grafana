// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana
{
    /// <summary>
    /// * [Official documentation](https://grafana.com/docs/grafana/latest/administration/team-management/)
    /// * [HTTP API](https://grafana.com/docs/grafana/latest/developers/http_api/team/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Grafana = Pulumiverse.Grafana;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var viewer = new Grafana.User("viewer", new()
    ///     {
    ///         Email = "viewer@example.com",
    ///         Login = "viewer",
    ///         Password = "my-password",
    ///     });
    /// 
    ///     var test_team = new Grafana.Team("test-team", new()
    ///     {
    ///         Email = "teamemail@example.com",
    ///         Members = new[]
    ///         {
    ///             viewer.Email,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import grafana:index/team:Team team_name {{team_id}} # To use the default provider org
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import grafana:index/team:Team team_name {{org_id}}:{{team_id}} # When "org_id" is set on the resource
    /// ```
    /// </summary>
    [GrafanaResourceType("grafana:index/team:Team")]
    public partial class Team : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An email address for the team.
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// Ignores team members that have been added to team by [Team
        /// Sync](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-team-sync/). Team Sync can be
        /// provisioned using [grafana_team_external_group
        /// resource](https://registry.terraform.io/providers/grafana/grafana/latest/docs/resources/team_external_group).
        /// </summary>
        [Output("ignoreExternallySyncedMembers")]
        public Output<bool?> IgnoreExternallySyncedMembers { get; private set; } = null!;

        /// <summary>
        /// A set of email addresses corresponding to users who should be given membership to the team. Note: users specified here
        /// must already exist in Grafana.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The display name for the Grafana team created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        [Output("preferences")]
        public Output<Outputs.TeamPreferences?> Preferences { get; private set; } = null!;

        /// <summary>
        /// The team id assigned to this team by Grafana.
        /// </summary>
        [Output("teamId")]
        public Output<int> TeamId { get; private set; } = null!;

        /// <summary>
        /// Sync external auth provider groups with this Grafana team. Only available in Grafana Enterprise. * [Official
        /// documentation](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-team-sync/) * [HTTP
        /// API](https://grafana.com/docs/grafana/latest/developers/http_api/team_sync/)
        /// </summary>
        [Output("teamSync")]
        public Output<Outputs.TeamTeamSync?> TeamSync { get; private set; } = null!;


        /// <summary>
        /// Create a Team resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Team(string name, TeamArgs? args = null, CustomResourceOptions? options = null)
            : base("grafana:index/team:Team", name, args ?? new TeamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Team(string name, Input<string> id, TeamState? state = null, CustomResourceOptions? options = null)
            : base("grafana:index/team:Team", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Team resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Team Get(string name, Input<string> id, TeamState? state = null, CustomResourceOptions? options = null)
        {
            return new Team(name, id, state, options);
        }
    }

    public sealed class TeamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An email address for the team.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Ignores team members that have been added to team by [Team
        /// Sync](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-team-sync/). Team Sync can be
        /// provisioned using [grafana_team_external_group
        /// resource](https://registry.terraform.io/providers/grafana/grafana/latest/docs/resources/team_external_group).
        /// </summary>
        [Input("ignoreExternallySyncedMembers")]
        public Input<bool>? IgnoreExternallySyncedMembers { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A set of email addresses corresponding to users who should be given membership to the team. Note: users specified here
        /// must already exist in Grafana.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The display name for the Grafana team created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("preferences")]
        public Input<Inputs.TeamPreferencesArgs>? Preferences { get; set; }

        /// <summary>
        /// Sync external auth provider groups with this Grafana team. Only available in Grafana Enterprise. * [Official
        /// documentation](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-team-sync/) * [HTTP
        /// API](https://grafana.com/docs/grafana/latest/developers/http_api/team_sync/)
        /// </summary>
        [Input("teamSync")]
        public Input<Inputs.TeamTeamSyncArgs>? TeamSync { get; set; }

        public TeamArgs()
        {
        }
        public static new TeamArgs Empty => new TeamArgs();
    }

    public sealed class TeamState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An email address for the team.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Ignores team members that have been added to team by [Team
        /// Sync](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-team-sync/). Team Sync can be
        /// provisioned using [grafana_team_external_group
        /// resource](https://registry.terraform.io/providers/grafana/grafana/latest/docs/resources/team_external_group).
        /// </summary>
        [Input("ignoreExternallySyncedMembers")]
        public Input<bool>? IgnoreExternallySyncedMembers { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A set of email addresses corresponding to users who should be given membership to the team. Note: users specified here
        /// must already exist in Grafana.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The display name for the Grafana team created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Organization ID. If not set, the Org ID defined in the provider block will be used.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("preferences")]
        public Input<Inputs.TeamPreferencesGetArgs>? Preferences { get; set; }

        /// <summary>
        /// The team id assigned to this team by Grafana.
        /// </summary>
        [Input("teamId")]
        public Input<int>? TeamId { get; set; }

        /// <summary>
        /// Sync external auth provider groups with this Grafana team. Only available in Grafana Enterprise. * [Official
        /// documentation](https://grafana.com/docs/grafana/latest/setup-grafana/configure-security/configure-team-sync/) * [HTTP
        /// API](https://grafana.com/docs/grafana/latest/developers/http_api/team_sync/)
        /// </summary>
        [Input("teamSync")]
        public Input<Inputs.TeamTeamSyncGetArgs>? TeamSync { get; set; }

        public TeamState()
        {
        }
        public static new TeamState Empty => new TeamState();
    }
}
