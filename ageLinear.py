# To add a new cell, type '# %%'
# To add a new markdown cell, type '# %% [markdown]'
# %%
import pandas as pd
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker
import numpy


# %%
def minCoef(age):
    age = age.values
    n = len(age)//2     # длина выборки
    p = n-1             # число отсчетов
    N = n*n
    k = [0.0 for i in range(n)]
    a = [0.0 for i in range(n)]
    R = [0.0 for i in range(n)]

    x = []
    for i in range(n):
        for j in range(n):
            x.append(age[i+j])
    # coefUl = []
    #
    # for i in range(n):
    #     c = []
    #     for j in range(n):
    #         flag = 0
    #         for l in range(n*n-abs(i-j)):
    #             flag += sys[l-i]*sys[l-j]
    #         flag /= n*n
    #         c.append(flag)
    #     coefUl.append(c)

    for i in range(n):
        flag = 0.0
        for h in range(N - i):
            flag += x[h]*x[h+i]
        R[i] = flag/N

    # R = coefUl[0]
    # print(R, len(R))
    E = R[0]
    l = 0

    for l in range(n):
        if l == 0:
            # E = R[0]
            continue

        for i in range(l):
            if i == 0:
                continue
            k[l] += a[i] * R[l - i]
        k[l] = (k[l] - R[l]) / E
        a[l] = -k[l]

        for j in range(l):
            if j == 0:
                continue
            a[j] += k[l] * a[i - j]
        E = E * (1 - k[l] * k[l])

    # print(a)
    ans = 0
    for i in range(n):
        ans += a[i]*x[n+i]
    print(ans)
    print(a)


# %%
# получение списка президентов
file = 'test.xlsx'
xl = pd.ExcelFile(file)
df = xl.parse('Лист1')  # список президентов
# print(df1)
# for age in df['ageIn'] :
#    print(age)
minCoef(df['ageIn'])


# %%
fig, ax = plt.subplots()

ax.set_title("Президенты USA")
ax.xaxis.set_major_locator(ticker.MultipleLocator(25))
ax.xaxis.set_minor_locator(ticker.MultipleLocator(5))

ax.yaxis.set_major_locator(ticker.MultipleLocator(5))
ax.yaxis.set_minor_locator(ticker.MultipleLocator(1))

ax.grid(which='major', color='k')
ax.minorticks_on()
ax.grid(which='minor', color='gray', linestyle=':')

plt.ylabel('Возраст')
plt.xlabel('Год')
plt.scatter(df["birth"], df["ageIn"], c='black', alpha=0.7)

fig.set_figwidth(16)
fig.set_figheight(8)

plt.show()
