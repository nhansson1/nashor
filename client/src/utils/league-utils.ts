import queues from "../static/queues.json";
export const getQueueNameById = (queueId: number): string => {
    const queue = queues.find(queue => queue.queueId === queueId);

    return queue?.description ?? "unknown mode";
};
