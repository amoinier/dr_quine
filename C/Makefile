# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: amoinier <amoinier@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2015/12/07 12:02:38 by amoinier          #+#    #+#              #
#    Updated: 2017/08/08 14:22:01 by amoinier         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

RED =		\033[31m
GRE =		\033[32m
YEL =		\033[33m
BLU =		\033[34m
CYA =		\033[36m
STD =		\033[39m

FLAG =		-Wall -Wextra -Werror

all: colleen grace sully

colleen:
	@echo "  ${BLU}+ Compilation program:${STD} $@"
	gcc $(FLAG) ./Colleen/Colleen.c -o ./Colleen/Colleen

grace:
	@echo "  ${BLU}+ Compilation program:${STD} $@"
	gcc $(FLAG) ./Grace/Grace.c -o ./Grace/Grace

sully:
	@echo "  ${BLU}+ Compilation program:${STD} $@"
	gcc $(FLAG) ./Sully/Sully.c -o ./Sully/Sully

clean:
	rm -f ./Colleen/Colleen ./Grace/Grace ./Sully/Sully

fclean: clean

.PHONY: all colleen grace sully clean fclean re

re: fclean all
