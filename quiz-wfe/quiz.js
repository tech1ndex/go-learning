//Javascript placeholder
const question = document.querySelector('#question');
const choices = Array.from(document.querySelectorAll('.choice.text'));
const progressText = document.querySelector('#progressText');
const scoreText = document.querySelector('#score');
const progressBarFull = document.querySelector('#progressBarFull');

let currentQuestion = {}
let acceptingAnswers = true
let score = 0
let questionCounter = 0
let availableQuestion = []

let questions = [
    {
        question: 'What is the gcloud command to set default zone for compute engine server using gcloud cli?',
        choice1: 'gcloud config set compute/zone us-east-1',
        choice2: 'gcloud config configurations set compute/zone us-east-1a',
        choice3: 'gcloud config set compute/zone us-east1-a',
        choice4: 'gcloud defaults set compute/zone us-east-1',
        answer: 3
    },
    {
        question: 'As a cloud engineer, you have been asked to upgrade the free trial of your account and rename it to a production-inventory system. You are getting permission denied error while making the changes. Which of the following permissions will solve the problem?',
        choice1: 'Billing.accounts.update',
        choice2: 'Billing.account.upgrade',
        choice3: 'Billing.account.update',
        choice4: 'Billing.accounts.upgrade',
        answer: 1
    }
]

const SCORE_POINTS = 100
const MAX_QUESTIONS = 2

startQuiz = () => 0 {
score = 0
availableQuestions = [...questions]
getNewQuestion() 
}

getNewQuestion = () => {
    if(availableQuestions.length === 0 || questionsCounter > MAX_QUESTIONS) {
        localStorage.setItem('mostRecentScore', score)

        return window.location.assign('/end.html')
    }

    questionCounter++
    progressText.innerText = `Question ${questionCounter} of ${MAX_QUESTIONS}`
    progressBarFull.style.width = `${(questionCounter/MAX_QUESTIONS) * 100}%`

    const questionsIndex = Math.floor(Math.random() * availableQuestions.length)
    currentQuestion = availableQuestions[questionsIndex]
    question.innerText = currentQuestion.question

    choices.forEach(choice => {
        const number = choice.dataset['number']
        choice.innerText = currentQuestion['choice' + number]
    })

    availableQuestions.splice(questionsIndex, 1)
    acceptingAnswers = true
}

choices.forEach(choice => {
    choice.addEventListener('click', e => {
        if(!acceptingAnswers) return

        acceptingAnswers = false
        const selectedChoice = e.target
        const selectedAnswer = selectedChoice.dataset['number']

        let classToApply = selectedAnswer == currentQuestion.answer ? 'correct' :
        'incorrect'

        if(classToApply === 'correct'){
            incrementScore(SCORE_POINTS)
        }

        selectedChoice.parentElement.classList.add(classToApply)

        setTimeout(() => {
            selectedChoice.parentElement.classList.remove(classToApply)
            getNewQuestion()
        
        }, 1000)
    })
})