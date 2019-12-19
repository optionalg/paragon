// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/kcarretto/paragon/ent/migrate"

	"github.com/kcarretto/paragon/ent/credential"
	"github.com/kcarretto/paragon/ent/job"
	"github.com/kcarretto/paragon/ent/jobtemplate"
	"github.com/kcarretto/paragon/ent/tag"
	"github.com/kcarretto/paragon/ent/target"
	"github.com/kcarretto/paragon/ent/task"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Credential is the client for interacting with the Credential builders.
	Credential *CredentialClient
	// Job is the client for interacting with the Job builders.
	Job *JobClient
	// JobTemplate is the client for interacting with the JobTemplate builders.
	JobTemplate *JobTemplateClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// Target is the client for interacting with the Target builders.
	Target *TargetClient
	// Task is the client for interacting with the Task builders.
	Task *TaskClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config:      c,
		Schema:      migrate.NewSchema(c.driver),
		Credential:  NewCredentialClient(c),
		Job:         NewJobClient(c),
		JobTemplate: NewJobTemplateClient(c),
		Tag:         NewTagClient(c),
		Target:      NewTargetClient(c),
		Task:        NewTaskClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil

	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config:      cfg,
		Credential:  NewCredentialClient(cfg),
		Job:         NewJobClient(cfg),
		JobTemplate: NewJobTemplateClient(cfg),
		Tag:         NewTagClient(cfg),
		Target:      NewTargetClient(cfg),
		Task:        NewTaskClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Credential.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config:      cfg,
		Schema:      migrate.NewSchema(cfg.driver),
		Credential:  NewCredentialClient(cfg),
		Job:         NewJobClient(cfg),
		JobTemplate: NewJobTemplateClient(cfg),
		Tag:         NewTagClient(cfg),
		Target:      NewTargetClient(cfg),
		Task:        NewTaskClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// CredentialClient is a client for the Credential schema.
type CredentialClient struct {
	config
}

// NewCredentialClient returns a client for the Credential from the given config.
func NewCredentialClient(c config) *CredentialClient {
	return &CredentialClient{config: c}
}

// Create returns a create builder for Credential.
func (c *CredentialClient) Create() *CredentialCreate {
	return &CredentialCreate{config: c.config}
}

// Update returns an update builder for Credential.
func (c *CredentialClient) Update() *CredentialUpdate {
	return &CredentialUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *CredentialClient) UpdateOne(cr *Credential) *CredentialUpdateOne {
	return c.UpdateOneID(cr.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *CredentialClient) UpdateOneID(id int) *CredentialUpdateOne {
	return &CredentialUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Credential.
func (c *CredentialClient) Delete() *CredentialDelete {
	return &CredentialDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CredentialClient) DeleteOne(cr *Credential) *CredentialDeleteOne {
	return c.DeleteOneID(cr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CredentialClient) DeleteOneID(id int) *CredentialDeleteOne {
	return &CredentialDeleteOne{c.Delete().Where(credential.ID(id))}
}

// Create returns a query builder for Credential.
func (c *CredentialClient) Query() *CredentialQuery {
	return &CredentialQuery{config: c.config}
}

// Get returns a Credential entity by its id.
func (c *CredentialClient) Get(ctx context.Context, id int) (*Credential, error) {
	return c.Query().Where(credential.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CredentialClient) GetX(ctx context.Context, id int) *Credential {
	cr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return cr
}

// JobClient is a client for the Job schema.
type JobClient struct {
	config
}

// NewJobClient returns a client for the Job from the given config.
func NewJobClient(c config) *JobClient {
	return &JobClient{config: c}
}

// Create returns a create builder for Job.
func (c *JobClient) Create() *JobCreate {
	return &JobCreate{config: c.config}
}

// Update returns an update builder for Job.
func (c *JobClient) Update() *JobUpdate {
	return &JobUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *JobClient) UpdateOne(j *Job) *JobUpdateOne {
	return c.UpdateOneID(j.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *JobClient) UpdateOneID(id int) *JobUpdateOne {
	return &JobUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Job.
func (c *JobClient) Delete() *JobDelete {
	return &JobDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *JobClient) DeleteOne(j *Job) *JobDeleteOne {
	return c.DeleteOneID(j.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *JobClient) DeleteOneID(id int) *JobDeleteOne {
	return &JobDeleteOne{c.Delete().Where(job.ID(id))}
}

// Create returns a query builder for Job.
func (c *JobClient) Query() *JobQuery {
	return &JobQuery{config: c.config}
}

// Get returns a Job entity by its id.
func (c *JobClient) Get(ctx context.Context, id int) (*Job, error) {
	return c.Query().Where(job.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *JobClient) GetX(ctx context.Context, id int) *Job {
	j, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return j
}

// QueryTasks queries the tasks edge of a Job.
func (c *JobClient) QueryTasks(j *Job) *TaskQuery {
	query := &TaskQuery{config: c.config}
	id := j.ID
	query.sql = sql.Select().From(sql.Table(task.Table)).
		Where(sql.EQ(job.TasksColumn, id))

	return query
}

// QueryTags queries the tags edge of a Job.
func (c *JobClient) QueryTags(j *Job) *TagQuery {
	query := &TagQuery{config: c.config}
	id := j.ID
	t1 := sql.Table(tag.Table)
	t2 := sql.Table(job.Table)
	t3 := sql.Table(job.TagsTable)
	t4 := sql.Select(t3.C(job.TagsPrimaryKey[1])).
		From(t3).
		Join(t2).
		On(t3.C(job.TagsPrimaryKey[0]), t2.C(job.FieldID)).
		Where(sql.EQ(t2.C(job.FieldID), id))
	query.sql = sql.Select().
		From(t1).
		Join(t4).
		On(t1.C(tag.FieldID), t4.C(job.TagsPrimaryKey[1]))

	return query
}

// QueryTemplate queries the template edge of a Job.
func (c *JobClient) QueryTemplate(j *Job) *JobTemplateQuery {
	query := &JobTemplateQuery{config: c.config}
	id := j.ID
	t1 := sql.Table(jobtemplate.Table)
	t2 := sql.Select(job.TemplateColumn).
		From(sql.Table(job.TemplateTable)).
		Where(sql.EQ(job.FieldID, id))
	query.sql = sql.Select().From(t1).Join(t2).On(t1.C(jobtemplate.FieldID), t2.C(job.TemplateColumn))

	return query
}

// JobTemplateClient is a client for the JobTemplate schema.
type JobTemplateClient struct {
	config
}

// NewJobTemplateClient returns a client for the JobTemplate from the given config.
func NewJobTemplateClient(c config) *JobTemplateClient {
	return &JobTemplateClient{config: c}
}

// Create returns a create builder for JobTemplate.
func (c *JobTemplateClient) Create() *JobTemplateCreate {
	return &JobTemplateCreate{config: c.config}
}

// Update returns an update builder for JobTemplate.
func (c *JobTemplateClient) Update() *JobTemplateUpdate {
	return &JobTemplateUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *JobTemplateClient) UpdateOne(jt *JobTemplate) *JobTemplateUpdateOne {
	return c.UpdateOneID(jt.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *JobTemplateClient) UpdateOneID(id int) *JobTemplateUpdateOne {
	return &JobTemplateUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for JobTemplate.
func (c *JobTemplateClient) Delete() *JobTemplateDelete {
	return &JobTemplateDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *JobTemplateClient) DeleteOne(jt *JobTemplate) *JobTemplateDeleteOne {
	return c.DeleteOneID(jt.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *JobTemplateClient) DeleteOneID(id int) *JobTemplateDeleteOne {
	return &JobTemplateDeleteOne{c.Delete().Where(jobtemplate.ID(id))}
}

// Create returns a query builder for JobTemplate.
func (c *JobTemplateClient) Query() *JobTemplateQuery {
	return &JobTemplateQuery{config: c.config}
}

// Get returns a JobTemplate entity by its id.
func (c *JobTemplateClient) Get(ctx context.Context, id int) (*JobTemplate, error) {
	return c.Query().Where(jobtemplate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *JobTemplateClient) GetX(ctx context.Context, id int) *JobTemplate {
	jt, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return jt
}

// QueryJobs queries the jobs edge of a JobTemplate.
func (c *JobTemplateClient) QueryJobs(jt *JobTemplate) *JobQuery {
	query := &JobQuery{config: c.config}
	id := jt.ID
	query.sql = sql.Select().From(sql.Table(job.Table)).
		Where(sql.EQ(jobtemplate.JobsColumn, id))

	return query
}

// QueryTags queries the tags edge of a JobTemplate.
func (c *JobTemplateClient) QueryTags(jt *JobTemplate) *TagQuery {
	query := &TagQuery{config: c.config}
	id := jt.ID
	t1 := sql.Table(tag.Table)
	t2 := sql.Table(jobtemplate.Table)
	t3 := sql.Table(jobtemplate.TagsTable)
	t4 := sql.Select(t3.C(jobtemplate.TagsPrimaryKey[1])).
		From(t3).
		Join(t2).
		On(t3.C(jobtemplate.TagsPrimaryKey[0]), t2.C(jobtemplate.FieldID)).
		Where(sql.EQ(t2.C(jobtemplate.FieldID), id))
	query.sql = sql.Select().
		From(t1).
		Join(t4).
		On(t1.C(tag.FieldID), t4.C(jobtemplate.TagsPrimaryKey[1]))

	return query
}

// TagClient is a client for the Tag schema.
type TagClient struct {
	config
}

// NewTagClient returns a client for the Tag from the given config.
func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

// Create returns a create builder for Tag.
func (c *TagClient) Create() *TagCreate {
	return &TagCreate{config: c.config}
}

// Update returns an update builder for Tag.
func (c *TagClient) Update() *TagUpdate {
	return &TagUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagClient) UpdateOne(t *Tag) *TagUpdateOne {
	return c.UpdateOneID(t.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *TagClient) UpdateOneID(id int) *TagUpdateOne {
	return &TagUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Tag.
func (c *TagClient) Delete() *TagDelete {
	return &TagDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TagClient) DeleteOne(t *Tag) *TagDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TagClient) DeleteOneID(id int) *TagDeleteOne {
	return &TagDeleteOne{c.Delete().Where(tag.ID(id))}
}

// Create returns a query builder for Tag.
func (c *TagClient) Query() *TagQuery {
	return &TagQuery{config: c.config}
}

// Get returns a Tag entity by its id.
func (c *TagClient) Get(ctx context.Context, id int) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id int) *Tag {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryTargets queries the targets edge of a Tag.
func (c *TagClient) QueryTargets(t *Tag) *TargetQuery {
	query := &TargetQuery{config: c.config}
	id := t.ID
	t1 := sql.Table(target.Table)
	t2 := sql.Table(tag.Table)
	t3 := sql.Table(tag.TargetsTable)
	t4 := sql.Select(t3.C(tag.TargetsPrimaryKey[0])).
		From(t3).
		Join(t2).
		On(t3.C(tag.TargetsPrimaryKey[1]), t2.C(tag.FieldID)).
		Where(sql.EQ(t2.C(tag.FieldID), id))
	query.sql = sql.Select().
		From(t1).
		Join(t4).
		On(t1.C(target.FieldID), t4.C(tag.TargetsPrimaryKey[0]))

	return query
}

// QueryTasks queries the tasks edge of a Tag.
func (c *TagClient) QueryTasks(t *Tag) *TaskQuery {
	query := &TaskQuery{config: c.config}
	id := t.ID
	t1 := sql.Table(task.Table)
	t2 := sql.Table(tag.Table)
	t3 := sql.Table(tag.TasksTable)
	t4 := sql.Select(t3.C(tag.TasksPrimaryKey[0])).
		From(t3).
		Join(t2).
		On(t3.C(tag.TasksPrimaryKey[1]), t2.C(tag.FieldID)).
		Where(sql.EQ(t2.C(tag.FieldID), id))
	query.sql = sql.Select().
		From(t1).
		Join(t4).
		On(t1.C(task.FieldID), t4.C(tag.TasksPrimaryKey[0]))

	return query
}

// QueryJobs queries the jobs edge of a Tag.
func (c *TagClient) QueryJobs(t *Tag) *JobQuery {
	query := &JobQuery{config: c.config}
	id := t.ID
	t1 := sql.Table(job.Table)
	t2 := sql.Table(tag.Table)
	t3 := sql.Table(tag.JobsTable)
	t4 := sql.Select(t3.C(tag.JobsPrimaryKey[0])).
		From(t3).
		Join(t2).
		On(t3.C(tag.JobsPrimaryKey[1]), t2.C(tag.FieldID)).
		Where(sql.EQ(t2.C(tag.FieldID), id))
	query.sql = sql.Select().
		From(t1).
		Join(t4).
		On(t1.C(job.FieldID), t4.C(tag.JobsPrimaryKey[0]))

	return query
}

// QueryJobTemplates queries the job_templates edge of a Tag.
func (c *TagClient) QueryJobTemplates(t *Tag) *JobTemplateQuery {
	query := &JobTemplateQuery{config: c.config}
	id := t.ID
	t1 := sql.Table(jobtemplate.Table)
	t2 := sql.Table(tag.Table)
	t3 := sql.Table(tag.JobTemplatesTable)
	t4 := sql.Select(t3.C(tag.JobTemplatesPrimaryKey[0])).
		From(t3).
		Join(t2).
		On(t3.C(tag.JobTemplatesPrimaryKey[1]), t2.C(tag.FieldID)).
		Where(sql.EQ(t2.C(tag.FieldID), id))
	query.sql = sql.Select().
		From(t1).
		Join(t4).
		On(t1.C(jobtemplate.FieldID), t4.C(tag.JobTemplatesPrimaryKey[0]))

	return query
}

// TargetClient is a client for the Target schema.
type TargetClient struct {
	config
}

// NewTargetClient returns a client for the Target from the given config.
func NewTargetClient(c config) *TargetClient {
	return &TargetClient{config: c}
}

// Create returns a create builder for Target.
func (c *TargetClient) Create() *TargetCreate {
	return &TargetCreate{config: c.config}
}

// Update returns an update builder for Target.
func (c *TargetClient) Update() *TargetUpdate {
	return &TargetUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *TargetClient) UpdateOne(t *Target) *TargetUpdateOne {
	return c.UpdateOneID(t.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *TargetClient) UpdateOneID(id int) *TargetUpdateOne {
	return &TargetUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Target.
func (c *TargetClient) Delete() *TargetDelete {
	return &TargetDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TargetClient) DeleteOne(t *Target) *TargetDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TargetClient) DeleteOneID(id int) *TargetDeleteOne {
	return &TargetDeleteOne{c.Delete().Where(target.ID(id))}
}

// Create returns a query builder for Target.
func (c *TargetClient) Query() *TargetQuery {
	return &TargetQuery{config: c.config}
}

// Get returns a Target entity by its id.
func (c *TargetClient) Get(ctx context.Context, id int) (*Target, error) {
	return c.Query().Where(target.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TargetClient) GetX(ctx context.Context, id int) *Target {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryTasks queries the tasks edge of a Target.
func (c *TargetClient) QueryTasks(t *Target) *TaskQuery {
	query := &TaskQuery{config: c.config}
	id := t.ID
	query.sql = sql.Select().From(sql.Table(task.Table)).
		Where(sql.EQ(target.TasksColumn, id))

	return query
}

// QueryTags queries the tags edge of a Target.
func (c *TargetClient) QueryTags(t *Target) *TagQuery {
	query := &TagQuery{config: c.config}
	id := t.ID
	t1 := sql.Table(tag.Table)
	t2 := sql.Table(target.Table)
	t3 := sql.Table(target.TagsTable)
	t4 := sql.Select(t3.C(target.TagsPrimaryKey[1])).
		From(t3).
		Join(t2).
		On(t3.C(target.TagsPrimaryKey[0]), t2.C(target.FieldID)).
		Where(sql.EQ(t2.C(target.FieldID), id))
	query.sql = sql.Select().
		From(t1).
		Join(t4).
		On(t1.C(tag.FieldID), t4.C(target.TagsPrimaryKey[1]))

	return query
}

// QueryCredentials queries the credentials edge of a Target.
func (c *TargetClient) QueryCredentials(t *Target) *CredentialQuery {
	query := &CredentialQuery{config: c.config}
	id := t.ID
	query.sql = sql.Select().From(sql.Table(credential.Table)).
		Where(sql.EQ(target.CredentialsColumn, id))

	return query
}

// TaskClient is a client for the Task schema.
type TaskClient struct {
	config
}

// NewTaskClient returns a client for the Task from the given config.
func NewTaskClient(c config) *TaskClient {
	return &TaskClient{config: c}
}

// Create returns a create builder for Task.
func (c *TaskClient) Create() *TaskCreate {
	return &TaskCreate{config: c.config}
}

// Update returns an update builder for Task.
func (c *TaskClient) Update() *TaskUpdate {
	return &TaskUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *TaskClient) UpdateOne(t *Task) *TaskUpdateOne {
	return c.UpdateOneID(t.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *TaskClient) UpdateOneID(id int) *TaskUpdateOne {
	return &TaskUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Task.
func (c *TaskClient) Delete() *TaskDelete {
	return &TaskDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TaskClient) DeleteOne(t *Task) *TaskDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TaskClient) DeleteOneID(id int) *TaskDeleteOne {
	return &TaskDeleteOne{c.Delete().Where(task.ID(id))}
}

// Create returns a query builder for Task.
func (c *TaskClient) Query() *TaskQuery {
	return &TaskQuery{config: c.config}
}

// Get returns a Task entity by its id.
func (c *TaskClient) Get(ctx context.Context, id int) (*Task, error) {
	return c.Query().Where(task.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TaskClient) GetX(ctx context.Context, id int) *Task {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// QueryTags queries the tags edge of a Task.
func (c *TaskClient) QueryTags(t *Task) *TagQuery {
	query := &TagQuery{config: c.config}
	id := t.ID
	t1 := sql.Table(tag.Table)
	t2 := sql.Table(task.Table)
	t3 := sql.Table(task.TagsTable)
	t4 := sql.Select(t3.C(task.TagsPrimaryKey[1])).
		From(t3).
		Join(t2).
		On(t3.C(task.TagsPrimaryKey[0]), t2.C(task.FieldID)).
		Where(sql.EQ(t2.C(task.FieldID), id))
	query.sql = sql.Select().
		From(t1).
		Join(t4).
		On(t1.C(tag.FieldID), t4.C(task.TagsPrimaryKey[1]))

	return query
}

// QueryJob queries the job edge of a Task.
func (c *TaskClient) QueryJob(t *Task) *JobQuery {
	query := &JobQuery{config: c.config}
	id := t.ID
	t1 := sql.Table(job.Table)
	t2 := sql.Select(task.JobColumn).
		From(sql.Table(task.JobTable)).
		Where(sql.EQ(task.FieldID, id))
	query.sql = sql.Select().From(t1).Join(t2).On(t1.C(job.FieldID), t2.C(task.JobColumn))

	return query
}
