const generateScores = (scores: MetricInput): MetricOutput => {
    const perCapita = scores.emissions / scores.employees
    const co2Score = Math.floor(perCapita) as Score
    return { co2Score: co2Score }
}

export type MetricInput = {
    emissions: number,
    employees: number,
}

export type MetricOutput = {
    co2Score: Score
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