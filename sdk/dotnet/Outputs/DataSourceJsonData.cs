// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Grafana.Outputs
{

    [OutputType]
    public sealed class DataSourceJsonData
    {
        /// <summary>
        /// (CloudWatch, Athena) The ARN of the role to be assumed by Grafana when using the CloudWatch or Athena data source.
        /// </summary>
        public readonly string? AssumeRoleArn;
        /// <summary>
        /// (CloudWatch, Athena) The authentication type used to access the data source.
        /// </summary>
        public readonly string? AuthType;
        /// <summary>
        /// (Stackdriver) The authentication type: `jwt` or `gce`.
        /// </summary>
        public readonly string? AuthenticationType;
        /// <summary>
        /// (Athena) Athena catalog.
        /// </summary>
        public readonly string? Catalog;
        /// <summary>
        /// (Stackdriver) Service account email address.
        /// </summary>
        public readonly string? ClientEmail;
        /// <summary>
        /// (Azure Monitor) The service account client id.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// (Azure Monitor) The cloud name.
        /// </summary>
        public readonly string? CloudName;
        /// <summary>
        /// (MySQL, PostgreSQL, and MSSQL) Maximum amount of time in seconds a connection may be reused (Grafana v5.4+).
        /// </summary>
        public readonly int? ConnMaxLifetime;
        /// <summary>
        /// (CloudWatch) A comma-separated list of custom namespaces to be queried by the CloudWatch data source.
        /// </summary>
        public readonly string? CustomMetricsNamespaces;
        /// <summary>
        /// (Athena) Name of the database within the catalog.
        /// </summary>
        public readonly string? Database;
        /// <summary>
        /// (InfluxDB) The default bucket for the data source.
        /// </summary>
        public readonly string? DefaultBucket;
        /// <summary>
        /// (Stackdriver) The default project for the data source.
        /// </summary>
        public readonly string? DefaultProject;
        /// <summary>
        /// (CloudWatch, Athena) The default region for the data source.
        /// </summary>
        public readonly string? DefaultRegion;
        /// <summary>
        /// (Loki) See https://grafana.com/docs/grafana/latest/datasources/loki/#derived-fields
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceJsonDataDerivedField> DerivedFields;
        /// <summary>
        /// (MSSQL) Connection SSL encryption handling: 'disable', 'false' or 'true'.
        /// </summary>
        public readonly string? Encrypt;
        /// <summary>
        /// (Elasticsearch) Elasticsearch semantic version (Grafana v8.0+).
        /// </summary>
        public readonly string? EsVersion;
        /// <summary>
        /// (CloudWatch, Athena) If you are assuming a role in another account, that has been created with an external ID, specify the external ID here.
        /// </summary>
        public readonly string? ExternalId;
        /// <summary>
        /// (Github) Github URL
        /// </summary>
        public readonly string? GithubUrl;
        /// <summary>
        /// (Graphite) Graphite version.
        /// </summary>
        public readonly string? GraphiteVersion;
        /// <summary>
        /// (Prometheus) HTTP method to use for making requests.
        /// </summary>
        public readonly string? HttpMethod;
        /// <summary>
        /// (Alertmanager) Implementation of Alertmanager. Either 'cortex' or 'prometheus'
        /// </summary>
        public readonly string? Implementation;
        /// <summary>
        /// (Elasticsearch) Index date time format. nil(No Pattern), 'Hourly', 'Daily', 'Weekly', 'Monthly' or 'Yearly'.
        /// </summary>
        public readonly string? Interval;
        /// <summary>
        /// (Elasticsearch) Which field should be used to indicate the priority of the log message.
        /// </summary>
        public readonly string? LogLevelField;
        /// <summary>
        /// (Elasticsearch) Which field should be used as the log message.
        /// </summary>
        public readonly string? LogMessageField;
        /// <summary>
        /// (Prometheus) Manage alerts.
        /// </summary>
        public readonly bool? ManageAlerts;
        /// <summary>
        /// (Elasticsearch) Maximum number of concurrent shard requests.
        /// </summary>
        public readonly int? MaxConcurrentShardRequests;
        /// <summary>
        /// (MySQL, PostgreSQL and MSSQL) Maximum number of connections in the idle connection pool (Grafana v5.4+).
        /// </summary>
        public readonly int? MaxIdleConns;
        /// <summary>
        /// (Loki) Upper limit for the number of log lines returned by Loki
        /// </summary>
        public readonly int? MaxLines;
        /// <summary>
        /// (MySQL, PostgreSQL and MSSQL) Maximum number of open connections to the database (Grafana v5.4+).
        /// </summary>
        public readonly int? MaxOpenConns;
        /// <summary>
        /// (Sentry) Organization slug.
        /// </summary>
        public readonly string? OrgSlug;
        /// <summary>
        /// (InfluxDB) An organization is a workspace for a group of users. All dashboards, tasks, buckets, members, etc., belong to an organization.
        /// </summary>
        public readonly string? Organization;
        /// <summary>
        /// (Athena) AWS S3 bucket to store execution outputs. If not specified, the default query result location from the Workgroup configuration will be used.
        /// </summary>
        public readonly string? OutputLocation;
        /// <summary>
        /// (PostgreSQL) Postgres version as a number (903/904/905/906/1000) meaning v9.3, v9.4, etc.
        /// </summary>
        public readonly int? PostgresVersion;
        /// <summary>
        /// (CloudWatch, Athena) The credentials profile name to use when authentication type is set as 'Credentials file'.
        /// </summary>
        public readonly string? Profile;
        /// <summary>
        /// (Prometheus) Timeout for queries made to the Prometheus data source in seconds.
        /// </summary>
        public readonly string? QueryTimeout;
        /// <summary>
        /// (Elasticsearch and Prometheus) Specifies the ARN of an IAM role to assume.
        /// </summary>
        public readonly string? Sigv4AssumeRoleArn;
        /// <summary>
        /// (Elasticsearch and Prometheus) Enable usage of SigV4.
        /// </summary>
        public readonly bool? Sigv4Auth;
        /// <summary>
        /// (Elasticsearch and Prometheus) The Sigv4 authentication provider to use: 'default', 'credentials' or 'keys' (AMG: 'workspace-iam-role').
        /// </summary>
        public readonly string? Sigv4AuthType;
        /// <summary>
        /// (Elasticsearch and Prometheus) When assuming a role in another account use this external ID.
        /// </summary>
        public readonly string? Sigv4ExternalId;
        /// <summary>
        /// (Elasticsearch and Prometheus) Credentials profile name, leave blank for default.
        /// </summary>
        public readonly string? Sigv4Profile;
        /// <summary>
        /// (Elasticsearch and Prometheus) AWS region to use for Sigv4.
        /// </summary>
        public readonly string? Sigv4Region;
        /// <summary>
        /// (PostgreSQL) SSLmode. 'disable', 'require', 'verify-ca' or 'verify-full'.
        /// </summary>
        public readonly string? SslMode;
        /// <summary>
        /// (Azure Monitor) The subscription id
        /// </summary>
        public readonly string? SubscriptionId;
        /// <summary>
        /// (Azure Monitor) Service account tenant ID.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// (Elasticsearch) Which field that should be used as timestamp.
        /// </summary>
        public readonly string? TimeField;
        /// <summary>
        /// (Prometheus, Elasticsearch, InfluxDB, MySQL, PostgreSQL, and MSSQL) Lowest interval/step value that should be used for this data source. Sometimes called "Scrape Interval" in the Grafana UI.
        /// </summary>
        public readonly string? TimeInterval;
        /// <summary>
        /// (PostgreSQL) Enable usage of TimescaleDB extension.
        /// </summary>
        public readonly bool? Timescaledb;
        /// <summary>
        /// (All) Enable TLS authentication using client cert configured in secure json data.
        /// </summary>
        public readonly bool? TlsAuth;
        /// <summary>
        /// (All) Enable TLS authentication using CA cert.
        /// </summary>
        public readonly bool? TlsAuthWithCaCert;
        /// <summary>
        /// (All) SSL Certificate configuration, either by ‘file-path’ or ‘file-content’.
        /// </summary>
        public readonly string? TlsConfigurationMethod;
        /// <summary>
        /// (All) Controls whether a client verifies the server’s certificate chain and host name.
        /// </summary>
        public readonly bool? TlsSkipVerify;
        /// <summary>
        /// (Stackdriver) The token URI used, provided in the service account key.
        /// </summary>
        public readonly string? TokenUri;
        /// <summary>
        /// (Cloudwatch) The X-Ray datasource uid to associate to this Cloudwatch datasource.
        /// </summary>
        public readonly string? TracingDatasourceUid;
        /// <summary>
        /// (OpenTSDB) Resolution.
        /// </summary>
        public readonly int? TsdbResolution;
        /// <summary>
        /// (OpenTSDB) Version.
        /// </summary>
        public readonly int? TsdbVersion;
        /// <summary>
        /// (InfluxDB) InfluxQL or Flux.
        /// </summary>
        public readonly string? Version;
        /// <summary>
        /// (Athena) Workgroup to use.
        /// </summary>
        public readonly string? Workgroup;
        /// <summary>
        /// (Elasticsearch) Enable X-Pack support.
        /// </summary>
        public readonly bool? XpackEnabled;

        [OutputConstructor]
        private DataSourceJsonData(
            string? assumeRoleArn,

            string? authType,

            string? authenticationType,

            string? catalog,

            string? clientEmail,

            string? clientId,

            string? cloudName,

            int? connMaxLifetime,

            string? customMetricsNamespaces,

            string? database,

            string? defaultBucket,

            string? defaultProject,

            string? defaultRegion,

            ImmutableArray<Outputs.DataSourceJsonDataDerivedField> derivedFields,

            string? encrypt,

            string? esVersion,

            string? externalId,

            string? githubUrl,

            string? graphiteVersion,

            string? httpMethod,

            string? implementation,

            string? interval,

            string? logLevelField,

            string? logMessageField,

            bool? manageAlerts,

            int? maxConcurrentShardRequests,

            int? maxIdleConns,

            int? maxLines,

            int? maxOpenConns,

            string? orgSlug,

            string? organization,

            string? outputLocation,

            int? postgresVersion,

            string? profile,

            string? queryTimeout,

            string? sigv4AssumeRoleArn,

            bool? sigv4Auth,

            string? sigv4AuthType,

            string? sigv4ExternalId,

            string? sigv4Profile,

            string? sigv4Region,

            string? sslMode,

            string? subscriptionId,

            string? tenantId,

            string? timeField,

            string? timeInterval,

            bool? timescaledb,

            bool? tlsAuth,

            bool? tlsAuthWithCaCert,

            string? tlsConfigurationMethod,

            bool? tlsSkipVerify,

            string? tokenUri,

            string? tracingDatasourceUid,

            int? tsdbResolution,

            int? tsdbVersion,

            string? version,

            string? workgroup,

            bool? xpackEnabled)
        {
            AssumeRoleArn = assumeRoleArn;
            AuthType = authType;
            AuthenticationType = authenticationType;
            Catalog = catalog;
            ClientEmail = clientEmail;
            ClientId = clientId;
            CloudName = cloudName;
            ConnMaxLifetime = connMaxLifetime;
            CustomMetricsNamespaces = customMetricsNamespaces;
            Database = database;
            DefaultBucket = defaultBucket;
            DefaultProject = defaultProject;
            DefaultRegion = defaultRegion;
            DerivedFields = derivedFields;
            Encrypt = encrypt;
            EsVersion = esVersion;
            ExternalId = externalId;
            GithubUrl = githubUrl;
            GraphiteVersion = graphiteVersion;
            HttpMethod = httpMethod;
            Implementation = implementation;
            Interval = interval;
            LogLevelField = logLevelField;
            LogMessageField = logMessageField;
            ManageAlerts = manageAlerts;
            MaxConcurrentShardRequests = maxConcurrentShardRequests;
            MaxIdleConns = maxIdleConns;
            MaxLines = maxLines;
            MaxOpenConns = maxOpenConns;
            OrgSlug = orgSlug;
            Organization = organization;
            OutputLocation = outputLocation;
            PostgresVersion = postgresVersion;
            Profile = profile;
            QueryTimeout = queryTimeout;
            Sigv4AssumeRoleArn = sigv4AssumeRoleArn;
            Sigv4Auth = sigv4Auth;
            Sigv4AuthType = sigv4AuthType;
            Sigv4ExternalId = sigv4ExternalId;
            Sigv4Profile = sigv4Profile;
            Sigv4Region = sigv4Region;
            SslMode = sslMode;
            SubscriptionId = subscriptionId;
            TenantId = tenantId;
            TimeField = timeField;
            TimeInterval = timeInterval;
            Timescaledb = timescaledb;
            TlsAuth = tlsAuth;
            TlsAuthWithCaCert = tlsAuthWithCaCert;
            TlsConfigurationMethod = tlsConfigurationMethod;
            TlsSkipVerify = tlsSkipVerify;
            TokenUri = tokenUri;
            TracingDatasourceUid = tracingDatasourceUid;
            TsdbResolution = tsdbResolution;
            TsdbVersion = tsdbVersion;
            Version = version;
            Workgroup = workgroup;
            XpackEnabled = xpackEnabled;
        }
    }
}
