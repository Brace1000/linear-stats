Project Overview

linear-stats is a command-line application written in Go that reads numerical data from a file, calculates the Linear Regression Line and Pearson Correlation Coefficient, and outputs the results. The project is designed to help users learn about statistical and probability calculations.
Features

    Reads data from a specified file
    Calculates the Linear Regression Line (y = mx + b)
    Calculates the Pearson Correlation Coefficient
    Outputs results with specified precision

Requirements

    Go 1.16 or higher

Installation

Clone the repository:

git clone https://learn.zone01kisumu.ke/git/bobaigwa/linear-stats
cd linear-stats

Usage

To run the program, use the following command:


go run . data.txt

Data File Format

The data file should contain comma-separated numerical values, one line per entry. Example:

189, 113, 121, 114, 145, 110
132, 119, 137, 115, 142, 120

Output Format

The program outputs the Linear Regression Line and Pearson Correlation Coefficient in the following format:

Linear Regression Line: y = <slope>x + <intercept>
Pearson Correlation Coefficient: <value>

Precision

    Linear Regression Line slope and intercept are printed with 6 decimal places.
    Pearson Correlation Coefficient is printed with 10 decimal places.



Testing

The program can be tested by running it with various data files and comparing the output to expected results.

    Generate a random data set and run the program:

    go run . data.txt

    Verify the output format and precision:
        Check that the Linear Regression Line is in the format y = <value>x + <value>.
        Ensure the values on the Linear Regression Line contain 6 decimal places.
        Check that the Pearson Correlation Coefficient contains 10 decimal places.

    Compare results:

    Run the program multiple times with different data sets to ensure consistent and accurate results.

Example Data

Below is an example of data to use for testing:

1267
1329
1354
1303
1299
1358

Contributing

If you wish to contribute to the project, please fork the repository and submit a pull request with your changes. For major changes, please open an issue first to discuss what you would like to change.
License
The project is licensed under the MIT License.