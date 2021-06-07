const SPEC_REPO = "datadog-api-spec"

module.exports.post_approved_pr_status_check = async (github, context) => {
    const pr_num = context.payload.pull_request.head.ref.split("/")[2]
    const {data: pr} = await github.pulls.get({
        owner: context.repo.owner,
        repo: SPEC_REPO,
        pull_number: pr_num,
    });
    let status = "pending"
    if (context.eventName === "pull_request_review" && context.payload.action === "submitted") {
        if (context.payload.review.state === "APPROVED") {
            status = "success"
        } else {
            status = "failure"
        }
    }
    await github.repos.createCommitStatus({
        owner: context.repo.owner,
        repo: SPEC_REPO,
        sha: pr.head.sha,
        state: status,
        target_url: `https://github.com/${context.repo.owner}/${context.repo.repo}/pull/${context.payload.pull_request.number}`,
        description: `PR on ${context.repo.repo} has been approved`,
        context: `${context.repo.repo}_pr_approved`
    });
}
