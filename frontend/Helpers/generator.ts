const generateScores = (scores: MetricInput): MetricOutput => {
    const output: MetricOutput = {
        actionabilityScore: Score.ZERO,
        awarenessScore: Score.ZERO,
        co2Score: Score.ZERO,
        energyScore: Score.ZERO,
        recyclingScore: Score.ZERO,
        waterScore: Score.ZERO,

    }
    const perCapita = scores.emissions / scores.employees
    const co2Score = Math.floor(perCapita) as Score
    return { ...output, co2Score: co2Score }
}

export type MetricInput = {
    emissions: number,
    employees: number,
}

export type MetricOutput = {
    co2Score: Score,
    waterScore: Score,
    energyScore: Score,
    recyclingScore: Score,
    awarenessScore: Score,
    actionabilityScore: Score,
}

export enum Score {
    ZERO,
    ONE,
    TWO,
    THREE,
    FOUR,
    FIVE,
    SIX,
    SEVEN,
    EIGHT,
    NINE,
    TEN
}