# Task

See feature overview here: `.github/overview.md`

For this task you are helping me to implement addition operator.

## UI Inputs:
- **Operators:** radio button: 
  - Addition
  - Subtraction
  - Multiplication
  - Division
- **Number of Questions:** text input (integer)
- **Number of Operands:** 2 to 3
- **Number of Digits for Each Operand:** 
    - Operand 1: text input (integer)
    - Operand 2: text input (integer)
    - ... according to the number of operands
- **Submit Button**

## Implementation Steps:
1. Create a form with the specified UI inputs.
2. Implement logic to handle the selected operator and generate questions accordingly.
3. Validate user inputs to ensure they meet the specified criteria.
4. Display the generated questions to the user.

## Retouching steps:

I have got a good start. I can see that when I input the information, the addition excercises are generated.
However, we need some improvement as below:

1. Remove duplications in the  `Generated Addition Problems`
2. Limit the number of operands to 3
3. Change the list to an ordered list
4. Add a print button and process printing
5. numOperands defaults to 2
6. pkg/handler/submit.go: please introduce a FormData struct to make it cleaner






