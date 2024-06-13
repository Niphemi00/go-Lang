# import pandas as pd
# import numpy as np
# # always do research 
# # METHODS OF PANDA DATA FRAme CREAIION
# # 1. DICTIONARY OF LIST
# data1 =  {
#   "Name" : ["Amy White", "Jack Stewart", "Richard Lauderdale", "Sara Jhonson"],
#   "Age" : [50, 53, 35, 43],
#   "Working Experience (Yrs)" : [5, 8, 3, 10]
# } 
# df1 = pd.DataFrame(data1)
# # 2. LIST OF DICTIONARIES
# data2 = [
#   {"Name" : "Amy White", "Age" : 50, "Working Experience(Yrs)." : 5},
#   {"Name" : "Jack Stewart", "Age" : 53, "Working Experience(Yrs)." : 8},
#   {"Name" : "Richard Lauderdale", "Age" : 35, "Working Experience(Yrs)." : 3},
#   {"Name" : "Sara Jhonson", "Age" : 43, "Working Experience(Yrs)." : 10}
# ]
# df2 = pd.DataFrame(data2)
# # 3. USING PANDAS SERIES
# series_balance= pd.Series([1000, 8459, 89394])
# series_name = pd.Series(["Joshua", "liam", "Ade"])
# series_age = pd.Series([22, 13, 45])
# data3 = {
#   "Name" : series_name,
#   "Balance" : series_balance,
#   "Age" : series_age
# }
# df3 = pd.DataFrame(data3)
# # 4. LIST OF LISTS
# data4 = [['Product A', 5000], ['Product B', 3000], ['Product C', 2300]]
# df4 = pd.DataFrame(data4)
# df4.columns = ['Products Name', 'Products price']


# # 5. PROFESSIONAL'S WAY
# df5 = pd.DataFrame(
#   data = [['Product A', 5000], ['Product B', 3000], ['Product C', 2300]],
#   columns = ['Products Name', 'Products price'],
#   index= ['A', 'B', 'C']
# )
# # OUTPUT RESULTS
# # print(df1)
# # print(df2)
# # print(df3)
# # print(df4)
# print(f'the shape is a {df5.shape} matrix')

def currency_converter():
    exchange_rates = {
        'USD': 1.0,       
        'EUR': 0.85,       
        'JPY': 110.0,      
        'GBP': 0.75,       
        'CAD': 1.25,
        'NGN':1500.0  
    }

    print("Available currencies: USD, EUR, JPY, GBP, CAD, NGN")
    
    from_currency = input("Enter the currency you want to convert from: ").upper()
    to_currency = input("Enter the currency you want to convert to: ").upper()
    amount = float(input("Enter the amount you want to convert: "))

    if from_currency not in exchange_rates or to_currency not in exchange_rates:
        print("Invalid currency code. Please try again.")
        return

    amount_in_usd = amount / exchange_rates[from_currency]

    converted_amount = amount_in_usd * exchange_rates[to_currency]

    print(f"{amount} {from_currency} is equal to {converted_amount:.2f} {to_currency}")

if __name__ == "__main__":
    print(__name__)
    while True:
      currency_converter()
