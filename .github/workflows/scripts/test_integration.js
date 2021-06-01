module.exports.post_status_check = async (github, context, status) => {
    const pr_num = context.payload.pull_request.head.ref.split("/")[2]
    const {data: pr} = await github.pulls.get({
        owner: context.repo.owner,
        repo: "datadog-api-spec",
        pull_number: pr_num,
    });
    const { data: jobs } = await github.actions.listJobsForWorkflowRun({
        owner: context.repo.owner,
        repo: context.repo.repo,
        run_id: context.runId
    });
    await github.repos.createCommitStatus({
        owner: context.repo.owner,
        repo: "datadog-api-spec",
        sha: pr.head.sha,
        state: status,
        target_url: `https://github.com/DataDog/datadog-api-client-go/pull/973/checks?check_run_id=${jobs.jobs[0].id}`,
        description: `${context.repo.repo} integration tests`,
        context: `${context.repo.repo}_integration_tests`
    });

}
