import queues from "../static/queues.json";
import champions from "../static/champion.json"
export const getQueueNameById = (queueId: number): string => {
    const queue = queues.find(queue => queue.queueId === queueId);

    return queue?.description ?? "unknown mode";
};

export const getChampionKeyById = (championId: number): string => {
    let championKey = ""

    for (const champion of Object.values(champions.data)) {
        if (champion.key === championId.toString()) {
            championKey = champion.id;
            break;
        }
    }

    return championKey
}
