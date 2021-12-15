# daysUntil

This is a simple terminal tool for showing the time until a given date/time.

### Usage

Multiple dates can be passed in the command line. Acceptable formats are either:

    mm-dd-yyyy
    mm-dd-yyyy HH:mm   (24hour clock)


### Example

    daysUntil  "12-10-2021 13:00" "12-11-2021 13:00" "12-17-2021 17:00"

### Notes

- After a date has passed, it is removed from view
- A date without a time is calculated to the start of the day
- Times are local